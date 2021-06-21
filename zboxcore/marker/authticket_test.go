package marker

import (
	"fmt"
	"testing"

	"github.com/0chain/gosdk/core/zcncrypto"
	zclient "github.com/0chain/gosdk/zboxcore/client"
	"github.com/stretchr/testify/require"
)

const (
	mockClientID        = "mock client id"
	mockOwnerID         = "mock owner id"
	mockAllocationID    = "mock allocation id"
	mockFilePathHash    = "mock file path hash"
	mockFileName        = "mock file name"
	mockRefType         = "mock ref type"
	mockReEncryptionKey = "mock re-encryption key"
	mockSignature       = "abcdef"
	mockPrivateKey      = "62fc118369fb9dd1fa6065d4f8f765c52ac68ad5aced17a1e5c4f8b4301a9469b987071c14695caf340ea11560f5a3cb76ad1e709803a8b339826ab3964e470a"
	mockPublicKey       = "b987071c14695caf340ea11560f5a3cb76ad1e709803a8b339826ab3964e470a"
)

func TestAuthTicket_GetHashData(t *testing.T) {
	require := require.New(t)
	rm := &AuthTicket{
		ClientID:        mockClientID,
		OwnerID:         mockOwnerID,
		AllocationID:    mockAllocationID,
		FilePathHash:    mockFilePathHash,
		FileName:        mockFileName,
		RefType:         mockRefType,
		ReEncryptionKey: mockReEncryptionKey,
		Expiration:      10,
		Timestamp:       10,
	}
	got := rm.GetHashData()
	require.EqualValues(got, mockAllocationID+":"+mockClientID+":"+mockOwnerID+":"+mockFilePathHash+":"+mockFileName+":"+mockRefType+":"+mockReEncryptionKey+":10:10")
}

func TestAuthTicket_Sign(t *testing.T) {
	require := require.New(t)
	err := zclient.PopulateClient("{}", "ed25519")
	if err != nil {
		fmt.Println(err)
	}
	client := zclient.GetClient()
	client.Wallet = &zcncrypto.Wallet{
		ClientID:  mockClientID,
		ClientKey: mockPublicKey,
		Keys: []zcncrypto.KeyPair{
			{
				PublicKey:  mockPublicKey,
				PrivateKey: mockPrivateKey,
			},
		},
	}
	rm := &AuthTicket{
		ClientID:        mockClientID,
		OwnerID:         mockOwnerID,
		AllocationID:    mockAllocationID,
		FilePathHash:    mockFilePathHash,
		FileName:        mockFileName,
		RefType:         mockRefType,
		ReEncryptionKey: mockReEncryptionKey,
		Expiration:      10,
		Timestamp:       10,
	}
	err = rm.Sign()
	require.NoErrorf(err, "Unexpected error %v", err)
}
