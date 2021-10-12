package magmasc

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/0chain/gosdk/zmagmacore/errors"
	"github.com/0chain/gosdk/zmagmacore/magmasc/pb"
	"github.com/0chain/gosdk/zmagmacore/time"
)

type (
	// AccessPoint represents access point node stored in blockchain.
	AccessPoint struct {
		*pb.AccessPoint
	}
)

// NewAccessPoint returns a new constructed provider terms.
func NewAccessPoint() *AccessPoint {
	return &AccessPoint{AccessPoint: &pb.AccessPoint{Terms: &pb.Terms{Qos: &pb.QoS{}}}}
}

// Decode implements util.Serializable interface.
func (m *AccessPoint) Decode(blob []byte) error {
	accessPoint := NewAccessPoint()
	if err := json.Unmarshal(blob, accessPoint); err != nil {
		return errDecodeData.Wrap(err)
	}
	if err := accessPoint.Validate(); err != nil {
		return err
	}

	m.AccessPoint = accessPoint.AccessPoint
	m.Terms = accessPoint.Terms

	return nil
}

// Encode implements util.Serializable interface.
func (m *AccessPoint) Encode() []byte {
	blob, _ := json.Marshal(m)
	return blob
}

// GetType returns node type.
func (m *AccessPoint) GetType() string {
	return accessPointType
}

// Validate checks the AccessPoint for correctness.
// If it is not return errInvalidAccessPoint.
func (m *AccessPoint) Validate() (err error) {
	switch { // is invalid
	case m.AccessPoint == nil:
		return errors.New(errCodeBadRequest, "accessPoint is not present yet")

	case m.Terms == nil:
		err = errors.New(errCodeBadRequest, "terms at accessPoint is not present yet")

	case m.Id == "":
		err = errors.New(errCodeBadRequest, "accessPoint external id is required")

	case m.ProviderExtId == "":
		err = errors.New(errCodeBadRequest, "accessPoint provider external id is required")
	}
	if err != nil {
		return errInvalidAccessPoint.Wrap(err)
	}

	if err = m.TermsValidate(); err != nil {
		return errInvalidAccessPoint.Wrap(err)
	}

	return nil
}

// ReadYAML reads config yaml file from path.
func (m *AccessPoint) ReadYAML(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) { _ = f.Close() }(f)

	decoder := yaml.NewDecoder(f)

	m.AccessPoint = NewAccessPoint().AccessPoint
	return decoder.Decode(m.AccessPoint)
}

// TermsDecrease makes automatically Decrease access point's terms by config.
func (m *AccessPoint) TermsDecrease() *AccessPoint {
	m.Terms.Volume = 0 // the volume of terms must be zeroed

	if m.Terms.ProlongDuration.Seconds != 0 {
		m.Terms.ExpiredAt.Seconds += m.Terms.ProlongDuration.Seconds // prolong expire of terms
	}

	if m.Terms.PriceAutoUpdate != 0 && m.Terms.Price > m.Terms.PriceAutoUpdate {
		m.Terms.Price -= m.Terms.PriceAutoUpdate // down the price
	}

	if m.Terms.QosAutoUpdate != nil {
		if m.Terms.QosAutoUpdate.UploadMbps != 0 {
			m.Terms.Qos.UploadMbps += m.Terms.QosAutoUpdate.UploadMbps // up the qos of upload mbps
		}

		if m.Terms.QosAutoUpdate.DownloadMbps != 0 {
			m.Terms.Qos.DownloadMbps += m.Terms.QosAutoUpdate.DownloadMbps // up the qos of download mbps
		}
	}

	return m
}

// TermsExpired returns if access point's terms already expired.
func (m *AccessPoint) TermsExpired() bool {
	return float64(m.Terms.ExpiredAt.Seconds) < time.Duration(time.Now()+TermsExpiredDuration).Seconds()
}

// TermsGetAmount returns calculated amount value of access point's terms.
func (m *AccessPoint) TermsGetAmount() (amount int64) {
	price := m.TermsGetPrice()
	if price > 0 {
		amount = price * m.TermsGetVolume()
		if minCost := m.TermsGetMinCost(); amount < minCost {
			amount = minCost
		}
	}

	return amount
}

// TermsGetMinCost returns calculated min cost value of access point's terms.
func (m *AccessPoint) TermsGetMinCost() (cost int64) {
	if m.Terms.MinCost > 0 {
		cost = int64(m.Terms.MinCost * billion)
	}

	return cost
}

// TermsGetPrice returns calculated price value of access point's terms.
// NOTE: the price value will be represented in token units per megabyte.
func (m *AccessPoint) TermsGetPrice() (price int64) {
	if m.Terms.Price > 0 {
		price = int64(m.Terms.Price * billion)
	}

	return price
}

// TermsGetVolume returns value of the provider terms volume.
// If the Volume is empty it will be calculated by the access point's terms.
func (m *AccessPoint) TermsGetVolume() int64 {
	if m.Terms.Volume == 0 {
		mbps := (m.Terms.Qos.UploadMbps + m.Terms.Qos.DownloadMbps) / octet // megabytes per second
		duration := float32(m.Terms.ExpiredAt.Seconds - int64(time.Now()))  // duration in seconds
		// rounded of bytes per second multiplied by duration in seconds
		m.Terms.Volume = int64(mbps * duration)
	}

	return m.Terms.Volume
}

// TermsIncrease makes automatically Increase access point's terms by config.
func (m *AccessPoint) TermsIncrease() *AccessPoint {
	m.Terms.Volume = 0 // the volume of terms must be zeroed

	if m.Terms.ProlongDuration.Seconds != 0 {
		m.Terms.ExpiredAt.Seconds += m.Terms.ProlongDuration.Seconds // prolong expire of terms
	}

	if m.Terms.PriceAutoUpdate != 0 {
		m.Terms.Price += m.Terms.PriceAutoUpdate // up the price
	}

	if m.Terms.QosAutoUpdate != nil {
		if m.Terms.QosAutoUpdate.UploadMbps != 0 && m.Terms.Qos.UploadMbps > m.Terms.QosAutoUpdate.UploadMbps {
			m.Terms.Qos.UploadMbps -= m.Terms.QosAutoUpdate.UploadMbps // down the qos of upload mbps
		}

		if m.Terms.QosAutoUpdate.DownloadMbps != 0 && m.Terms.Qos.DownloadMbps > m.Terms.QosAutoUpdate.DownloadMbps {
			m.Terms.Qos.DownloadMbps -= m.Terms.QosAutoUpdate.DownloadMbps // down the qos of download mbps
		}
	}

	return m
}

// TermsValidate checks access point's terms for correctness.
// If it is not return errInvalidTerms.
func (m *AccessPoint) TermsValidate() (err error) {
	switch { // is invalid
	case m.Terms == nil:
		err = errors.New(errCodeBadRequest, "terms is not present yet")

	case m.Terms.Qos == nil:
		err = errors.New(errCodeBadRequest, "invalid terms qos")

	case m.Terms.Qos.UploadMbps <= 0:
		err = errors.New(errCodeBadRequest, "invalid terms qos upload mbps")

	case m.Terms.Qos.DownloadMbps <= 0:
		err = errors.New(errCodeBadRequest, "invalid terms qos download mbps")

	case m.TermsExpired():
		now := time.NowTime().Add(TermsExpiredDuration).Format(time.RFC3339)
		err = errors.New(errCodeBadRequest, "expired at must be after "+now)

	default:
		return nil // is valid
	}

	return errInvalidTerms.Wrap(err)
}
