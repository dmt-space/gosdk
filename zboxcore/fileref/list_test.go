package fileref

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	mockAllocationRoot = "mock allocation root"
)

func TestListResult_GetDirTree(t *testing.T) {
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
			errMsg:  "invalid_list_path: Invalid list path. list was not for a directory",
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
			lr := &ListResult{
				AllocationRoot: mockAllocationRoot,
				Meta:           tt.parameters.Meta,
				Entities:       nil,
			}
			got, err := lr.GetDirTree(mockAllocationID)
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

func TestListResult_populateChildren(t *testing.T) {
	require := require.New(t)
	lr := &ListResult{
		AllocationRoot: mockAllocationRoot,
		Meta: map[string]interface{}{
			"type": "d",
		},
		Entities: []map[string]interface{}{
			{
				"type": "d",
			},
		},
	}
	err := lr.populateChildren(&Ref{
		Type:         "d",
		AllocationID: mockAllocationID,
	})
	require.NoErrorf(err, "unexpected error %v", err)
}
