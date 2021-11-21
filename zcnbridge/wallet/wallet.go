package wallet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/0chain/gosdk/core/zcncrypto"
	"github.com/0chain/gosdk/zcnbridge/common"
	"github.com/0chain/gosdk/zcnbridge/crypto"
	"github.com/0chain/gosdk/zcnbridge/errors"
	"github.com/0chain/gosdk/zcncore"
)

const (
	ZCNSCSmartContractAddress = "6dba10422e368813802877a85039d3985d96760ed844092319743fb3a76712e0"
	AddAuthorizerFunc         = "AddAuthorizer"
	DeleteAuthorizerFunc      = "DeleteAuthorizer"
	MintFunc                  = "mint"
	BurnFunc                  = "burn"
	// ConsensusThresh quorum required to reach consensus
	ConsensusThresh      = float64(70.0)
	BurnWzcnTicketPath   = "/v1/ether/burnticket/get"
	BurnNativeTicketPath = "/v1/0chain/burnticket/get"
)

type (
	// Wallet represents a wallet that stores keys and additional info.
	Wallet struct {
		ZCNWallet *zcncrypto.Wallet
	}
)

func AssignWallet(clientConfig string) (*Wallet, error) {
	w := &zcncrypto.Wallet{}
	err := json.Unmarshal([]byte(clientConfig), w)
	if err != nil {
		return nil, errors.Wrap("unmarshal", "error while unmarshalling the wallet", err)
	}

	return &Wallet{w}, nil
}

// CreateWallet creates initialized Wallet.
func CreateWallet(publicKey, privateKey []byte) *Wallet {
	var (
		publicKeyHex, privateKeyHex = hex.EncodeToString(publicKey), hex.EncodeToString(privateKey)
	)
	return &Wallet{
		ZCNWallet: &zcncrypto.Wallet{
			ClientID:  crypto.Hash(publicKey),
			ClientKey: publicKeyHex,
			Keys: []zcncrypto.KeyPair{
				{
					PublicKey:  publicKeyHex,
					PrivateKey: privateKeyHex,
				},
			},
			Version: zcncrypto.CryptoVersion,
		},
	}
}

// PublicKey returns the public key.
func (w *Wallet) PublicKey() string {
	return w.ZCNWallet.Keys[0].PublicKey
}

// ID returns the client id.
// NOTE: client id represents hex encoded SHA3-256 hash of the raw public key.
func (w *Wallet) ID() string {
	return w.ZCNWallet.ClientID
}

// StringJSON returns marshalled to JSON string Wallet.ZCNWallet.
func (w *Wallet) StringJSON() (string, error) {
	byt, err := json.Marshal(w.ZCNWallet)
	if err != nil {
		return "", err
	}

	return string(byt), err
}

// RegisterToMiners registers wallet to the miners by executing zcncore.RegisterToMiners.
func (w *Wallet) RegisterToMiners() error {
	const errCode = "register_wallet"

	walletStr, err := w.StringJSON()
	if err != nil {
		return errors.Wrap(errCode, "error while marshalling wallet", err)
	}
	if err := zcncore.SetWalletInfo(walletStr, false); err != nil {
		return errors.Wrap(errCode, "error while init wallet", err)
	}

	wg := &sync.WaitGroup{}
	statusBar := &common.ZCNStatus{Wg: wg}
	wg.Add(1)
	err = zcncore.RegisterToMiners(w.ZCNWallet, statusBar)
	if err != nil {
		return errors.Wrap(errCode, "error while init wallet", err)
	}
	wg.Wait()
	if statusBar.Success {
		fmt.Println("wallet registered")
	} else {
		return errors.Wrap(errCode, "wallet registration failed "+statusBar.ErrMsg, err)
	}

	return nil
}