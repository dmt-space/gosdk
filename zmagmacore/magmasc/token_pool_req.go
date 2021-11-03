package magmasc

import (
	"encoding/json"

	"github.com/0chain/gosdk/core/util"
	"github.com/0chain/gosdk/zmagmacore/errors"
	"github.com/0chain/gosdk/zmagmacore/magmasc/pb"
)

type (
	// TokenPoolReq represents lock pool request implementation.
	TokenPoolReq struct {
		*pb.TokenPoolReq
	}
)

var (
	// Make sure TokenPoolReq implements Serializable interface.
	_ util.Serializable = (*TokenPoolReq)(nil)
)

// NewTokenPoolReq creates initialized TokenPoolReq.
func NewTokenPoolReq() *TokenPoolReq {
	return &TokenPoolReq{TokenPoolReq: &pb.TokenPoolReq{}}
}

// Decode implements util.Serializable interface.
func (m *TokenPoolReq) Decode(blob []byte) error {
	req := NewTokenPoolReq()
	if err := json.Unmarshal(blob, &req); err != nil {
		return ErrDecodeData.Wrap(err)
	}
	if err := req.Validate(); err != nil {
		return err
	}

	m.TokenPoolReq = req.TokenPoolReq

	return nil
}

// Encode implements util.Serializable interface.
func (m *TokenPoolReq) Encode() []byte {
	blob, _ := json.Marshal(m)
	return blob
}

// PoolID implements PoolConfigurator interface.
func (m *TokenPoolReq) PoolID() string {
	return m.Id
}

// PoolHolderID implements PoolConfigurator interface.
func (m *TokenPoolReq) PoolHolderID() string {
	return Address
}

// PoolPayeeID implements PoolConfigurator interface.
func (m *TokenPoolReq) PoolPayeeID() string {
	return m.PayeeId
}

// Validate checks TokenPoolReq for correctness.
func (m *TokenPoolReq) Validate() (err error) {
	if m.Id == "" { // is invalid
		err = errors.New(ErrCodeBadRequest, "pool id is required")
	}

	return err
}
