package magmasc

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/0chain/gosdk/core/util"
	"github.com/0chain/gosdk/zmagmacore/errors"
	"github.com/0chain/gosdk/zmagmacore/magmasc/pb"
)

type (
	// Provider represents providers node stored in blockchain.
	Provider struct {
		*pb.Provider
	}
)

var (
	// Make sure Provider implements Serializable interface.
	_ util.Serializable = (*Provider)(nil)
)

// NewProvider creates initialized Provider.
func NewProvider() *Provider {
	return &Provider{Provider: &pb.Provider{}}
}

// Decode implements util.Serializable interface.
func (m *Provider) Decode(blob []byte) error {
	var provider Provider
	if err := json.Unmarshal(blob, &provider); err != nil {
		return ErrDecodeData.Wrap(err)
	}
	if err := provider.Validate(); err != nil {
		return err
	}

	m.Provider = provider.Provider

	return nil
}

// Encode implements util.Serializable interface.
func (m *Provider) Encode() []byte {
	blob, _ := json.Marshal(m)
	return blob
}

// ExternalID returns the external id of Provider node.
func (m *Provider) ExternalID() string {
	return m.ExtId
}

// Validate checks Provider for correctness.
// If it is not return errInvalidProvider.
func (m *Provider) Validate() (err error) {
	switch { // is invalid
	case m.Provider == nil:
		err = errors.New(ErrCodeBadRequest, "provider is not present yet")

	case m.ExtId == "":
		err = errors.New(ErrCodeBadRequest, "provider external id is required")

	case m.Host == "":
		err = errors.New(ErrCodeBadRequest, "provider host is required")

	default:
		return nil // is valid
	}

	return ErrInvalidProvider.Wrap(err)
}

// ReadYAML reads config yaml file from path.
func (m *Provider) ReadYAML(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) { _ = f.Close() }(f)

	decoder := yaml.NewDecoder(f)

	m.Provider = NewProvider().Provider
	return decoder.Decode(m.Provider)
}
