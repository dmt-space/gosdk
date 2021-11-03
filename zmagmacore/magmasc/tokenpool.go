package magmasc

import (
	"encoding/json"

	"github.com/0chain/gosdk/core/util"
	"github.com/0chain/gosdk/zmagmacore/magmasc/pb"
)

type (
	// TokenPool represents token pool implementation.
	TokenPool struct {
		*pb.TokenPool
	}
)

var (
	// Make sure tokenPool implements Serializable interface.
	_ util.Serializable = (*TokenPool)(nil)
)

// NewTokenPool creates initialized NewTokenPool.
func NewTokenPool() *TokenPool {
	return &TokenPool{TokenPool: &pb.TokenPool{Transfers: []*pb.TokenPoolTransfer{}}}
}

// Decode implements util.Serializable interface.
func (m *TokenPool) Decode(blob []byte) error {
	pool := NewTokenPool()
	if err := json.Unmarshal(blob, &pool); err != nil {
		return ErrDecodeData.Wrap(err)
	}

	m.TokenPool = pool.TokenPool

	return nil
}

// Encode implements util.Serializable interface.
func (m *TokenPool) Encode() []byte {
	blob, _ := json.Marshal(m)
	return blob
}
