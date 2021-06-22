package fileref

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReferencePath_GetRefFromObjectTree(t *testing.T) {
	type parameters struct {
		Meta          map[string]interface{}
		ExpectedValue interface{}
	}
	tests := []struct {
		name       string
		parameters parameters
	}{
		{
			name: "Test_Dir_Tree",
			parameters: parameters{
				Meta: map[string]interface{}{
					"type": "d",
				},
				ExpectedValue: &Ref{
					Type:         "d",
					AllocationID: mockAllocationID,
				},
			},
		},
		{
			name: "Test_Get_Ref_From_Object_Tree",
			parameters: parameters{
				Meta: map[string]interface{}{
					"type": "f",
				},
				ExpectedValue: &FileRef{
					Ref: Ref{
						Type:         "f",
						AllocationID: mockAllocationID,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			rp := &ReferencePath{
				Meta: tt.parameters.Meta,
				List: nil,
			}
			got, _ := rp.GetRefFromObjectTree(mockAllocationID)
			require.EqualValues(got, tt.parameters.ExpectedValue)
		})
	}
}

func TestReferencePath_GetDirTree(t *testing.T) {
	type parameters struct {
		Meta map[string]interface{}
	}
	tests := []struct {
		name       string
		parameters parameters
		wantErr    bool
		errMsg     string
	}{
		{
			name: "Test_Error_Invalid_List_Path",
			parameters: parameters{
				Meta: map[string]interface{}{
					"type": "f",
				},
			},
			wantErr: true,
			errMsg:  "invalid_ref_path: Invalid reference path. root was not a directory type",
		},
		{
			name: "Test_Success",
			parameters: parameters{
				Meta: map[string]interface{}{
					"type": "d",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			rp := &ReferencePath{
				Meta: tt.parameters.Meta,
				List: nil,
			}
			got, err := rp.GetDirTree(mockAllocationID)
			require.EqualValues(tt.wantErr, err != nil)
			if err != nil {
				require.EqualValues(tt.errMsg, err.Error())
				return
			}
			expectedValue := &Ref{
				Type:         "d",
				AllocationID: mockAllocationID,
			}
			require.EqualValues(got, expectedValue)
		})
	}
}

func TestReferencePath_populateChildren(t *testing.T) {
	require := require.New(t)
	rp := &ReferencePath{
		Meta: map[string]interface{}{
			"type": "d",
		},
		List: []*ReferencePath{
			{
				Meta: map[string]interface{}{
					"type": "d",
				},
			},
		},
	}
	err := rp.populateChildren(&Ref{
		Type:         "d",
		AllocationID: mockAllocationID,
	})
	require.NoErrorf(err, "unexpected error %v", err)
}
