package magmasc

import (
	"encoding/json"
	"github.com/0chain/gosdk/core/util"
)

type (
	// CommonRewardPool represents common reward pool wrapper implementation.
	CommonRewardPool struct {
		ID string `json:"id"`
	}
)

var (
	// Make sure common reward pool implements Serializable interface.
	_ util.Serializable = (*CommonRewardPool)(nil)
)

// newCommonRewPool returns a new constructed token pool.
func newCommonRewPool() *CommonRewardPool {
	return &CommonRewardPool{}
}

// Decode implements util.Serializable interface.
func (m *CommonRewardPool) Decode(blob []byte) error {
	var pool CommonRewardPool
	if err := json.Unmarshal(blob, &pool); err != nil {
		return errDecodeData.Wrap(err)
	}

	m.ID = pool.ID

	return nil
}

// Encode implements util.Serializable interface.
func (m *CommonRewardPool) Encode() []byte {
	blob, _ := json.Marshal(m)
	return blob
}
