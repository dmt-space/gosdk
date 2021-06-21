package marker

import (
	"fmt"
	"testing"

	"github.com/0chain/gosdk/core/encryption"
	"github.com/0chain/gosdk/core/zcncrypto"
	zclient "github.com/0chain/gosdk/zboxcore/client"
	"github.com/stretchr/testify/require"
)

func TestReadMarker_GetHash(t *testing.T) {
	require := require.New(t)
	rm := &ReadMarker{
		ClientID:        mockClientID,
		ClientPublicKey: mockPublicKey,
		BlobberID:       mockBlobberID,
		AllocationID:    mockAllocationID,
		OwnerID:         mockOwnerID,
		Timestamp:       10,
		ReadCounter:     10,
		Signature:       mockSignature,
	}
	got := rm.GetHash()
	require.EqualValues(got, encryption.Hash(mockAllocationID+":"+mockBlobberID+":"+mockClientID+":"+mockPublicKey+":"+mockOwnerID+":10:10"))
}

func TestReadMarker_Sign(t *testing.T) {
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
	rm := &ReadMarker{
		ClientID:        mockClientID,
		ClientPublicKey: mockPublicKey,
		BlobberID:       mockBlobberID,
		AllocationID:    mockAllocationID,
		OwnerID:         mockOwnerID,
		Timestamp:       10,
		ReadCounter:     10,
		Signature:       mockSignature,
	}
	err = rm.Sign()
	require.NoErrorf(err, "Unexpected error %v", err)
}
