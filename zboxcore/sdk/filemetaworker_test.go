package sdk

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/0chain/gosdk/zboxcore/client"
	"github.com/0chain/gosdk/zboxcore/fileref"
	"github.com/0chain/gosdk/zboxcore/mocks"
	"github.com/0chain/gosdk/zboxcore/zboxutil"
	"github.com/stretchr/testify/mock"
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
		additionalMock func(t *testing.T) (teardown func(t *testing.T))
		want           bool
		wantErr        bool
	}{
		{
			"Test_Error_New_HTTP_Failed",
			func(t *testing.T) (teardown func(t *testing.T)) {
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
			func(t *testing.T) (teardown func(t *testing.T)) {
				m := &mocks.HttpClient{}
				zboxutil.Client = m
				m.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(strings.NewReader("This is not JSON format")),
				}, nil)
				return nil
			},
			false,
			true,
		},
		{
			"Test_Success_With_Auth_Ticket",
			func(t *testing.T) (teardown func(t *testing.T)) {
				m := &mocks.HttpClient{}
				zboxutil.Client = m
				bodyString := `{"actual_file_hash":"03cfd743661f07975fa2f1220c5194cbaff48451","actual_file_size":4,"actual_thumbnail_hash":"","actual_thumbnail_size":0,"attributes":{},"collaborators":[],"commit_meta_txns":[],"content_hash":"3a52ce780950d4d969792a2559cd519d7ee8c727","created_at":"2021-03-17T08:15:36.137135Z","custom_meta":"","encrypted_key":"","hash":"f7e67105e7645b71dc8a6dc92d68d989026eb1dd2817f1b042153c9a7a4d7a13","lookup_hash":"c884abb32aa0357e2541b683f6e52bfab9143d33b968977cf6ba31b43e832697","merkle_root":"ea13052ab648c94a2fc001ce4f6f5f2d8bb699d4b69264b361c45324c88da744","mimetype":"application/octet-stream","name":"1.txt","num_of_blocks":1,"on_cloud":false,"path":"/1.txt","path_hash":"c884abb32aa0357e2541b683f6e52bfab9143d33b968977cf6ba31b43e832697","size":1,"thumbnail_hash":"","thumbnail_size":0,"type":"f","updated_at":"2021-03-17T08:15:36.137135Z"}`
				m.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(strings.NewReader(bodyString)),
				}, nil)
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
			func(t *testing.T) (teardown func(t *testing.T)) {
				m := &mocks.HttpClient{}
				zboxutil.Client = m
				bodyString := `{"actual_file_hash":"03cfd743661f07975fa2f1220c5194cbaff48451","actual_file_size":4,"actual_thumbnail_hash":"","actual_thumbnail_size":0,"attributes":{},"collaborators":[],"commit_meta_txns":[],"content_hash":"3a52ce780950d4d969792a2559cd519d7ee8c727","created_at":"2021-03-17T08:15:36.137135Z","custom_meta":"","encrypted_key":"","hash":"f7e67105e7645b71dc8a6dc92d68d989026eb1dd2817f1b042153c9a7a4d7a13","lookup_hash":"c884abb32aa0357e2541b683f6e52bfab9143d33b968977cf6ba31b43e832697","merkle_root":"ea13052ab648c94a2fc001ce4f6f5f2d8bb699d4b69264b361c45324c88da744","mimetype":"application/octet-stream","name":"1.txt","num_of_blocks":1,"on_cloud":false,"path":"/1.txt","path_hash":"c884abb32aa0357e2541b683f6e52bfab9143d33b968977cf6ba31b43e832697","size":1,"thumbnail_hash":"","thumbnail_size":0,"type":"f","updated_at":"2021-03-17T08:15:36.137135Z"}`
				m.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(strings.NewReader(bodyString)),
				}, nil)
				return nil
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			if additionalMock := tt.additionalMock; additionalMock != nil {
				if teardown := additionalMock(t); teardown != nil {
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
	defer cncl()
	tests := []struct {
		name           string
		additionalMock func(t *testing.T) (teardown func(t *testing.T))
		wantErr        bool
		wantFunc       func(require *require.Assertions, req *ListRequest, foundMask zboxutil.Uint128)
	}{
		{
			"Test_All_Success",
			func(t *testing.T) (teardown func(t *testing.T)) {
				m := &mocks.HttpClient{}
				zboxutil.Client = m
				bodyString := `{"actual_file_hash":"03cfd743661f07975fa2f1220c5194cbaff48451","actual_file_size":4,"actual_thumbnail_hash":"","actual_thumbnail_size":0,"attributes":{},"collaborators":[],"commit_meta_txns":[],"content_hash":"3a52ce780950d4d969792a2559cd519d7ee8c727","created_at":"2021-03-17T08:15:36.137135Z","custom_meta":"","encrypted_key":"","hash":"f7e67105e7645b71dc8a6dc92d68d989026eb1dd2817f1b042153c9a7a4d7a13","lookup_hash":"c884abb32aa0357e2541b683f6e52bfab9143d33b968977cf6ba31b43e832697","merkle_root":"ea13052ab648c94a2fc001ce4f6f5f2d8bb699d4b69264b361c45324c88da744","mimetype":"application/octet-stream","name":"1.txt","num_of_blocks":1,"on_cloud":false,"path":"/1.txt","path_hash":"c884abb32aa0357e2541b683f6e52bfab9143d33b968977cf6ba31b43e832697","size":1,"thumbnail_hash":"","thumbnail_size":0,"type":"f","updated_at":"2021-03-17T08:15:36.137135Z"}`
				m.On("Do", mock.AnythingOfType("*http.Request")).Run(func(args mock.Arguments) {
					for _, c := range m.ExpectedCalls {
						c.ReturnArguments = mock.Arguments{&http.Response{
							StatusCode: http.StatusOK,
							Body:       ioutil.NopCloser(strings.NewReader(bodyString)),
						}, nil}
					}
				})
				return nil
			},
			false,
			func(require *require.Assertions, req *ListRequest, foundMask zboxutil.Uint128) {
				require.NotNil(req)
				require.Equal(float32(3), req.consensus)
				require.Equal(zboxutil.NewUint128(0xf), foundMask, "found value must be same")
			},
		},
		{
			"Test_Index_3_Error",
			func(t *testing.T) (teardown func(t *testing.T)) {
				m := &mocks.HttpClient{}
				zboxutil.Client = m
				bodyString := `{"actual_file_hash":"03cfd743661f07975fa2f1220c5194cbaff48451","actual_file_size":4,"actual_thumbnail_hash":"","actual_thumbnail_size":0,"attributes":{},"collaborators":[],"commit_meta_txns":[],"content_hash":"3a52ce780950d4d969792a2559cd519d7ee8c727","created_at":"2021-03-17T08:15:36.137135Z","custom_meta":"","encrypted_key":"","hash":"f7e67105e7645b71dc8a6dc92d68d989026eb1dd2817f1b042153c9a7a4d7a13","lookup_hash":"c884abb32aa0357e2541b683f6e52bfab9143d33b968977cf6ba31b43e832697","merkle_root":"ea13052ab648c94a2fc001ce4f6f5f2d8bb699d4b69264b361c45324c88da744","mimetype":"application/octet-stream","name":"1.txt","num_of_blocks":1,"on_cloud":false,"path":"/1.txt","path_hash":"c884abb32aa0357e2541b683f6e52bfab9143d33b968977cf6ba31b43e832697","size":1,"thumbnail_hash":"","thumbnail_size":0,"type":"f","updated_at":"2021-03-17T08:15:36.137135Z"}`
				statusCode := http.StatusOK
				mockCall := m.On("Do", mock.MatchedBy(func(req *http.Request) bool { return req.Method == "POST" }))
				mockCall.RunFn = func(args mock.Arguments) {
					req := args[0].(*http.Request)
					url := req.URL.Host
					switch url {
					case strings.ReplaceAll(a.Blobbers[3].Baseurl, "http://", ""):
						statusCode = http.StatusBadRequest
					default:
						statusCode = http.StatusOK
					}
					mockCall.ReturnArguments = mock.Arguments{&http.Response{
						StatusCode: statusCode,
						Body:       ioutil.NopCloser(strings.NewReader(bodyString)),
					}, nil}
				}
				return nil
			},
			false,
			func(require *require.Assertions, req *ListRequest, foundMask zboxutil.Uint128) {
				require.NotNil(req)
				require.Equal(float32(3), req.consensus)
				require.Equal(zboxutil.NewUint128(0x7), foundMask, "found value must be same")
			},
		},
		{
			"Test_Error_File_Consensus_Not_Found",
			func(t *testing.T) (teardown func(t *testing.T)) {
				m := &mocks.HttpClient{}
				zboxutil.Client = m
				bodyString := `{"actual_file_hash":"03cfd743661f07975fa2f1220c5194cbaff48451","actual_file_size":4,"actual_thumbnail_hash":"","actual_thumbnail_size":0,"attributes":{},"collaborators":[],"commit_meta_txns":[],"content_hash":"3a52ce780950d4d969792a2559cd519d7ee8c727","created_at":"2021-03-17T08:15:36.137135Z","custom_meta":"","encrypted_key":"","hash":"f7e67105e7645b71dc8a6dc92d68d989026eb1dd2817f1b042153c9a7a4d7a13","lookup_hash":"c884abb32aa0357e2541b683f6e52bfab9143d33b968977cf6ba31b43e832697","merkle_root":"ea13052ab648c94a2fc001ce4f6f5f2d8bb699d4b69264b361c45324c88da744","mimetype":"application/octet-stream","name":"1.txt","num_of_blocks":1,"on_cloud":false,"path":"/1.txt","path_hash":"c884abb32aa0357e2541b683f6e52bfab9143d33b968977cf6ba31b43e832697","size":1,"thumbnail_hash":"","thumbnail_size":0,"type":"f","updated_at":"2021-03-17T08:15:36.137135Z"}`
				m.On("Do", mock.AnythingOfType("*http.Request")).Run(func(args mock.Arguments) {
					for _, c := range m.ExpectedCalls {
						c.ReturnArguments = mock.Arguments{&http.Response{
							StatusCode: http.StatusBadRequest,
							Body:       ioutil.NopCloser(strings.NewReader(bodyString)),
						}, nil}
					}
				})
				return nil
			},
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
				if teardown := additionalMock(t); teardown != nil {
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
