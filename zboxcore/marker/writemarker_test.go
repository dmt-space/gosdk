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
	mockAllocationRoot         = "mock allocation root"
	mockPreviousAllocationRoot = "mock previous allocation root"
)

func TestWriteMarker_GetHashData(t *testing.T) {
	require := require.New(t)
	wm := &WriteMarker{
		AllocationRoot:         mockAllocationRoot,
		PreviousAllocationRoot: mockPreviousAllocationRoot,
		AllocationID:           mockAllocationID,
		Size:                   10,
		BlobberID:              mockBlobberID,
		Timestamp:              10,
		ClientID:               mockClientID,
		Signature:              mockSignature,
	}
	got := wm.GetHashData()
	require.EqualValues(got, mockAllocationRoot+":"+mockPreviousAllocationRoot+":"+mockAllocationID+":"+mockBlobberID+":"+mockClientID+":10:10")
}

func TestWriteMarker_GetHash(t *testing.T) {
	require := require.New(t)
	wm := &WriteMarker{
		AllocationRoot:         mockAllocationRoot,
		PreviousAllocationRoot: mockPreviousAllocationRoot,
		AllocationID:           mockAllocationID,
		Size:                   10,
		BlobberID:              mockBlobberID,
		Timestamp:              10,
		ClientID:               mockClientID,
		Signature:              mockSignature,
	}
	got := wm.GetHash()
	require.EqualValues(got, encryption.Hash(mockAllocationRoot+":"+mockPreviousAllocationRoot+":"+mockAllocationID+":"+mockBlobberID+":"+mockClientID+":10:10"))
}

func TestWriteMarker_Sign(t *testing.T) {
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
	wm := &WriteMarker{
		AllocationRoot:         mockAllocationRoot,
		PreviousAllocationRoot: mockPreviousAllocationRoot,
		AllocationID:           mockAllocationID,
		Size:                   10,
		BlobberID:              mockBlobberID,
		Timestamp:              10,
		ClientID:               mockClientID,
		Signature:              mockSignature,
	}
	err = wm.Sign()
	require.NoErrorf(err, "Unexpected error %v", err)
}

func TestWriteMarker_VerifySignature(t *testing.T) {
	type parameters struct {
		clientPublicKey string
		signature       string
	}
	tests := []struct {
		name       string
		parameters parameters
		setup      func(*testing.T, string)
		wantErr    bool
		errMsg     string
	}{
		{
			name: "Test_Error_During_Verifying_Signature",
			parameters: parameters{
				clientPublicKey: "mock client key",
				signature:       "",
			},
			wantErr: true,
			errMsg:  "write_marker_validation_failed: Error during verifying signature. public key does not exists for verification",
		},
		{
			name: "Test_Error_Write_Marker_Signature_Is_Not_Valid",
			parameters: parameters{
				clientPublicKey: "b987071c14695caf340ea11560f5a3cb76ad1e709803a8b339826ab3964e470a",
				signature:       "abcdef",
			},
			wantErr: true,
			errMsg:  "write_marker_validation_failed: Write marker signature is not valid",
		},
		{
			name: "Test_No_Error",
			parameters: parameters{
				clientPublicKey: mockPublicKey,
				signature: func() string {
					signScheme := zcncrypto.NewSignatureScheme("ed25519")
					signature, err := signScheme.Sign(mockAllocationRoot + ":" + mockPreviousAllocationRoot + ":" + mockAllocationID + ":" + mockBlobberID + ":" + mockClientID + ":10:10")
					if err == nil {
						t.Fatalf("Sign passed without private key")
					}
					signScheme.SetPrivateKey(mockPrivateKey)
					signature, err = signScheme.Sign(encryption.Hash(mockAllocationRoot + ":" + mockPreviousAllocationRoot + ":" + mockAllocationID + ":" + mockBlobberID + ":" + mockClientID + ":10:10"))
					if err != nil {
						t.Fatalf("ed25519 signing failed")
					}
					return signature
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			err := zclient.PopulateClient("{}", "ed25519")
			if err != nil {
				fmt.Println(err)
			}
			client := zclient.GetClient()
			client.Wallet = &zcncrypto.Wallet{
				ClientID:  mockClientID,
				ClientKey: tt.parameters.clientPublicKey,
				Keys: []zcncrypto.KeyPair{
					{
						PublicKey:  mockPublicKey,
						PrivateKey: mockPrivateKey,
					},
				},
			}
			wm := &WriteMarker{
				AllocationRoot:         mockAllocationRoot,
				PreviousAllocationRoot: mockPreviousAllocationRoot,
				AllocationID:           mockAllocationID,
				Size:                   10,
				BlobberID:              mockBlobberID,
				Timestamp:              10,
				ClientID:               mockClientID,
				Signature:              tt.parameters.signature,
			}
			err = wm.VerifySignature(mockPublicKey)
			require.EqualValues(tt.wantErr, err != nil)
			if err != nil {
				require.EqualValues(tt.errMsg, err.Error())
				return
			}
			require.NoErrorf(err, "Unexpected error %v", err)
		})
	}
}
