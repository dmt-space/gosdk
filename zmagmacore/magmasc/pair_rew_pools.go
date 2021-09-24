package magmasc

import (
	"github.com/0chain/gosdk/zmagmacore/errors"
)

type (
	// PairRewardPool implementation of the access point-provider pair.
	PairRewardPool struct {
		AccessPointID string `json:"access_point_id,omitempty"`
		ProviderID    string `json:"provider_id,omitempty"`
	}
)

// Validate checks the PairRewardPool for correctness.
// If it is not return errInvalidPairRewardPool.
func (m *PairRewardPool) Validate() (err error) {
	switch { // is invalid
	case m.AccessPointID == "":
		err = errors.New(errCodeBadRequest, "access point id is required")

	case m.ProviderID == "":
		err = errors.New(errCodeBadRequest, "provider id in pair is required")

	default:
		return nil // is valid
	}

	return errInvalidPairRewardPool.Wrap(err)
}
