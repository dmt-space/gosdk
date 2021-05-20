package sdk

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"github.com/0chain/gosdk/zboxcore/client"
	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/zboxutil"
	"github.com/stretchr/testify/require"
)

const (
	filemetaWorkerTestDir = configDir + "/filemetaworker"
)

func TestListRequest_getFileMetaInfoFromBlobber(t *testing.T) {
	// setup mock sdk
	_, _, blobberMocks, closeFn := setupMockInitStorageSDK(t, configDir, 1)
	defer closeFn()
	// setup mock allocation
	a, cncl := setupMockAllocation(t, filemetaWorkerTestDir, blobberMocks)
	var blobbersResponseMock = func(t *testing.T, testcaseName string) (teardown func(t *testing.T)) {
		setupBlobberMockResponses(t, blobberMocks, filemetaWorkerTestDir+"/getFileMetaInfoFromBlobber", testcaseName)
		return nil
	}
	defer cncl()
	req := &ListRequest{
		allocationID:   a.ID,
		allocationTx:   a.Tx,
		blobbers:       a.Blobbers,
		remotefilepath: "/1.txt",
		Consensus: Consensus{
			consensusThresh: (float32(a.DataShards) * 100) / float32(a.DataShards+a.ParityShards),
			fullconsensus:   float32(a.DataShards + a.ParityShards),
		},
		ctx: context.Background(),
		wg:  &sync.WaitGroup{},
	}
	tests := []struct {
		name           string
		additionalMock func(t *testing.T, testCaseName string) (teardown func(t *testing.T))
		want           bool
		wantErr        bool
	}{
		{
			"Test_Error_New_HTTP_Failed",
			func(t *testing.T, testCaseName string) (teardown func(t *testing.T)) {
				url := a.Blobbers[0].Baseurl
				a.Blobbers[0].Baseurl = string([]byte{0x7f, 0, 0})
				return func(t *testing.T) {
					a.Blobbers[0].Baseurl = url
				}
			},
			false,
			true,
		},
		{
			"Test_HTTP_Response_Failed",
			nil,
			false,
			false,
		},
		{
			"Test_Error_HTTP_Response_Not_JSON_Format",
			blobbersResponseMock,
			false,
			true,
		},
		{
			"Test_Success_With_Auth_Ticket",
			func(t *testing.T, testCaseName string) (teardown func(t *testing.T)) {
				blobbersResponseMock(t, testCaseName)
				authTicket, err := a.GetAuthTicket("/1/txt", "1.txt", fileref.FILE, client.GetClientID(), "")
				require.NoErrorf(t, err, "unexpected get auth ticket error: %v", err)
				require.NotEmptyf(t, authTicket, "unexpected empty auth ticket")
				sEnc, err := base64.StdEncoding.DecodeString(authTicket)
				require.NoErrorf(t, err, "unexpected decode auth ticket error: %v", err)
				err = json.Unmarshal(sEnc, &req.authToken)
				require.NoErrorf(t, err, "unexpected error when marshaling auth ticket error: %v", err)
				return func(t *testing.T) {
					req.authToken = nil
				}
			},
			true,
			false,
		},
		{
			"Test_Success",
			blobbersResponseMock,
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			if additionalMock := tt.additionalMock; additionalMock != nil {
				if teardown := additionalMock(t, tt.name); teardown != nil {
					defer teardown(t)
				}
			}
			rspCh := make(chan *fileMetaResponse, 1)
			req.wg.Add(1)
			go req.getFileMetaInfoFromBlobber(req.blobbers[0], 0, rspCh)
			req.wg.Wait()
			resp := <-rspCh

			var expectedResult *fileref.FileRef
			parseFileContent(t, fmt.Sprintf("%v/%v/expected_result__Test_Success.json", filemetaWorkerTestDir, "getFileMetaInfoFromBlobber"), &expectedResult)
			if tt.wantErr {
				require.Error(resp.err, "expected error != nil")
				return
			}
			if !tt.want {
				require.NotEqual(expectedResult, resp.fileref)
				return
			}
			require.EqualValues(expectedResult, resp.fileref)
		})
	}
}

func TestListRequest_getFileConsensusFromBlobbers(t *testing.T) {
	// setup mock sdk
	_, _, blobberMocks, closeFn := setupMockInitStorageSDK(t, configDir, 4)
	defer closeFn()
	// setup mock allocation
	a, cncl := setupMockAllocation(t, filemetaWorkerTestDir, blobberMocks)
	var blobbersResponseMock = func(t *testing.T, testcaseName string) (teardown func(t *testing.T)) {
		setupBlobberMockResponses(t, blobberMocks, filemetaWorkerTestDir+"/getFileConsensusFromBlobbers", testcaseName)
		return nil
	}
	defer cncl()
	tests := []struct {
		name           string
		additionalMock func(t *testing.T, testCaseName string) (teardown func(t *testing.T))
		wantErr        bool
		wantFunc       func(require *require.Assertions, req *ListRequest, foundMask zboxutil.Uint128)
	}{
		{
			"Test_All_Success",
			blobbersResponseMock,
			false,
			func(require *require.Assertions, req *ListRequest, foundMask zboxutil.Uint128) {
				require.NotNil(req)
				require.Equal(float32(3), req.consensus)
				require.Equal(zboxutil.NewUint128(0xf), foundMask, "found value must be same")
			},
		},
		{
			"Test_Index_3_Error",
			blobbersResponseMock,
			false,
			func(require *require.Assertions, req *ListRequest, foundMask zboxutil.Uint128) {
				require.NotNil(req)
				require.Equal(float32(3), req.consensus)
				require.Equal(zboxutil.NewUint128(0x7), foundMask, "found value must be same")
			},
		},
		{
			"Test_Error_File_Consensus_Not_Found",
			blobbersResponseMock,
			true,
			func(require *require.Assertions, req *ListRequest, foundMask zboxutil.Uint128) {
				require.NotNil(req)
				require.Equal(float32(0), req.consensus)
				require.Equal(zboxutil.NewUint128(0x0), foundMask, "found value must be same")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			if additionalMock := tt.additionalMock; additionalMock != nil {
				if teardown := additionalMock(t, tt.name); teardown != nil {
					defer teardown(t)
				}
			}
			req := &ListRequest{
				allocationID:   a.ID,
				allocationTx:   a.Tx,
				blobbers:       a.Blobbers,
				remotefilepath: "/1.txt",
				ctx:            context.Background(),
				Consensus: Consensus{
					consensusThresh: (float32(a.DataShards) * 100) / float32(a.DataShards+a.ParityShards),
					fullconsensus:   float32(a.DataShards + a.ParityShards),
				},
			}

			foundMask, fileRef, _ := req.getFileConsensusFromBlobbers()
			var expectedResult *fileref.FileRef
			parseFileContent(t, fmt.Sprintf("%v/%v/expected_result__Test_Success.json", filemetaWorkerTestDir, "getFileConsensusFromBlobbers"), &expectedResult)
			if !tt.wantErr {
				require.EqualValues(expectedResult, fileRef)
			}
			tt.wantFunc(require, req, foundMask)
		})
	}
}
