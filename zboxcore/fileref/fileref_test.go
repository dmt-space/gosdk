package fileref

import (
	"encoding/json"
	"testing"

	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/core/encryption"
	"github.com/stretchr/testify/require"
)

const (
	mockType           = "f"
	mockName           = "mock name"
	mockAllocationID   = "mock allocation id"
	mockPath           = "mock path"
	mockHash           = "mock hash"
	mockNumBlocks      = 10
	mockPathHash       = "mock path hash"
	mockLookupHash     = "mock lookup hash"
	mockContentHash    = "mock content hash"
	mockMerkleRoot     = "mock merkle root"
	mockActualFileHash = "mock actual file hash"
)

func TestGetReferenceLookup(t *testing.T) {
	require := require.New(t)
	got := GetReferenceLookup(mockAllocationID, mockPath)
	require.EqualValues(got, encryption.Hash(mockAllocationID+":"+mockPath))
}

func TestRef_CalculateHash(t *testing.T) {
	type parameters struct {
		childrenLoaded bool
		Children       []RefEntity
	}
	tests := []struct {
		name         string
		parameters   parameters
		expectedHash string
	}{
		{
			name: "Test_Children_Array_Empty",
			parameters: parameters{
				Children:       nil,
				childrenLoaded: false,
			},
			expectedHash: "mock hash",
		},
		{
			name: "Test_Success",
			parameters: parameters{
				Children: []RefEntity{
					&Ref{
						Type:      "f",
						Path:      "mock path",
						Hash:      "mock hash",
						NumBlocks: 10,
						PathHash:  "mock path hash",
					},
				},
				childrenLoaded: true,
			},
			expectedHash: encryption.Hash("mock hash"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			r := &Ref{
				AllocationID:   mockAllocationID,
				Path:           mockPath,
				Size:           10,
				ActualSize:     10,
				Hash:           mockHash,
				NumBlocks:      mockNumBlocks,
				PathHash:       mockPathHash,
				LookupHash:     mockLookupHash,
				childrenLoaded: tt.parameters.childrenLoaded,
				Children:       tt.parameters.Children,
			}
			got := r.CalculateHash()
			require.EqualValues(got, tt.expectedHash)
		})
	}
}

func TestRef_AddChild(t *testing.T) {
	require := require.New(t)
	r := &Ref{
		AllocationID: mockAllocationID,
		Path:         mockPath,
		Size:         10,
		ActualSize:   10,
		Hash:         mockHash,
		NumBlocks:    mockNumBlocks,
		PathHash:     mockPathHash,
		LookupHash:   mockLookupHash,
		Children:     []RefEntity{},
	}
	child := []RefEntity{
		&Ref{
			Type:      "f",
			Path:      "mock path",
			Hash:      "mock hash",
			NumBlocks: 10,
			PathHash:  "mock path hash",
		},
	}
	r.AddChild(child[0])
	require.EqualValues(r.Children, child)
	require.EqualValues(r.childrenLoaded, true)
}

func TestRef_RemoveChild(t *testing.T) {
	require := require.New(t)
	r := &Ref{
		AllocationID: mockAllocationID,
		Path:         mockPath,
		Size:         10,
		ActualSize:   10,
		Hash:         mockHash,
		NumBlocks:    mockNumBlocks,
		PathHash:     mockPathHash,
		LookupHash:   mockLookupHash,
		Children: []RefEntity{
			&Ref{
				Type:      "f",
				Path:      "mock path",
				Hash:      "mock hash",
				NumBlocks: 10,
				PathHash:  "mock path hash",
			},
		},
	}
	r.RemoveChild(0)
	require.EqualValues(r.Children, []RefEntity{})
}

func TestFileRef_GetHashData(t *testing.T) {
	require := require.New(t)
	fr := &FileRef{
		Ref: Ref{
			AllocationID: mockAllocationID,
			Name:         mockName,
			Path:         mockPath,
			Size:         10,
			ActualSize:   10,
			Hash:         mockHash,
			NumBlocks:    mockNumBlocks,
			PathHash:     mockPathHash,
			LookupHash:   mockLookupHash,
			Type:         mockType,
		},
		ContentHash:    mockContentHash,
		MerkleRoot:     mockMerkleRoot,
		ActualFileHash: mockActualFileHash,
		ActualFileSize: 10,
		Attributes:     Attributes{WhoPaysForReads: common.WhoPays3rdParty},
	}
	got := fr.GetHashData()
	var attrs, _ = json.Marshal(&fr.Attributes)
	expectedValue := mockAllocationID + ":" + mockType + ":" + mockName + ":" + mockPath + ":10:" + mockContentHash + ":" + mockMerkleRoot + ":10:" + mockActualFileHash + ":" + string(attrs)
	require.EqualValues(got, expectedValue)
}

func TestFileRef_CalculateHash(t *testing.T) {
	require := require.New(t)
	fr := &FileRef{
		Ref: Ref{
			AllocationID: mockAllocationID,
			Name:         mockName,
			Path:         mockPath,
			Size:         10,
			ActualSize:   10,
			Hash:         mockHash,
			NumBlocks:    mockNumBlocks,
			PathHash:     mockPathHash,
			LookupHash:   mockLookupHash,
			Type:         mockType,
		},
		ContentHash:    mockContentHash,
		MerkleRoot:     mockMerkleRoot,
		ActualFileHash: mockActualFileHash,
		ActualFileSize: 10,
		Attributes:     Attributes{WhoPaysForReads: common.WhoPays3rdParty},
	}
	got := fr.CalculateHash()
	var attrs, _ = json.Marshal(&fr.Attributes)
	expectedValue := mockAllocationID + ":" + mockType + ":" + mockName + ":" + mockPath + ":10:" + mockContentHash + ":" + mockMerkleRoot + ":10:" + mockActualFileHash + ":" + string(attrs)
	require.EqualValues(got, encryption.Hash(expectedValue))
}
