package marker

import (
	"fmt"
	"testing"

	"github.com/0chain/gosdk/core/encryption"
	"github.com/0chain/gosdk/core/zcncrypto"
	zclient "github.com/0chain/gosdk/zboxcore/client"
	"github.com/stretchr/testify/require"
)

const (
	mockBlobberID   = "mock blobber id"
	mockFileRefHash = "mock file ref hash"
)

func TestDeleteToken_GetHash(t *testing.T) {
	require := require.New(t)
	dt := &DeleteToken{
		ClientID:     mockClientID,
		BlobberID:    mockBlobberID,
		AllocationID: mockAllocationID,
		FilePathHash: mockFilePathHash,
		FileRefHash:  mockFileRefHash,
		Signature:    mockSignature,
		Size:         10,
		Timestamp:    10,
	}
	got := dt.GetHash()
	fmt.Println(got)
	require.EqualValues(got, encryption.Hash(mockFileRefHash+":"+mockFilePathHash+":"+mockAllocationID+":"+mockBlobberID+":"+mockClientID+":10:10"))
}

func TestDeleteToken_Sign(t *testing.T) {
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
	dt := &DeleteToken{
		ClientID:     mockClientID,
		BlobberID:    mockBlobberID,
		AllocationID: mockAllocationID,
		FilePathHash: mockFilePathHash,
		FileRefHash:  mockFileRefHash,
		Signature:    mockSignature,
		Size:         10,
		Timestamp:    10,
	}
	err = dt.Sign()
	require.NoErrorf(err, "Unexpected error %v", err)
}
