package magmasc

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_TokenPoolReq_Decode(t *testing.T) {
	t.Parallel()

	tokenPoolReq := mockTokenPoolReq()
	blob, err := json.Marshal(tokenPoolReq)
	if err != nil {
		t.Fatalf("json.Marshal() error: %v | want: %v", err, nil)
	}

	tests := [2]struct {
		name  string
		blob  []byte
		want  *TokenPoolReq
		error bool
	}{
		{
			name:  "OK",
			blob:  blob,
			want:  tokenPoolReq,
			error: false,
		},
		{
			name:  "Decode_ERR",
			blob:  []byte(":"), // invalid json
			want:  &TokenPoolReq{},
			error: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := &TokenPoolReq{}
			if err := got.Decode(test.blob); (err != nil) != test.error {
				t.Errorf("Decode() error: %v | want: %v", err, test.error)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Decode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_TokenPoolReq_Encode(t *testing.T) {
	t.Parallel()

	tokenPoolReq := mockTokenPoolReq()
	blob, err := json.Marshal(tokenPoolReq)
	if err != nil {
		t.Fatalf("json.Marshal() error: %v | want: %v", err, nil)
	}

	tests := [1]struct {
		name         string
		tokenPoolReq *TokenPoolReq
		want         []byte
	}{
		{
			name:         "OK",
			tokenPoolReq: tokenPoolReq,
			want:         blob,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.tokenPoolReq.Encode(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Encode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_TokenPoolReq_Validate(t *testing.T) {
	t.Parallel()

	tokenPoolReq := mockTokenPoolReq()
	tokenPoolReqEmptyID := mockTokenPoolReq()
	tokenPoolReqEmptyID.Id = ""

	tests := [2]struct {
		name         string
		tokenPoolReq *TokenPoolReq
		error        bool
	}{
		{
			name:         "OK",
			tokenPoolReq: tokenPoolReq,
			error:        false,
		},
		{
			name:         "Empty_ID_ERR",
			tokenPoolReq: tokenPoolReqEmptyID,
			error:        true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if err := test.tokenPoolReq.Validate(); (err != nil) != test.error {
				t.Errorf("validate() error: %v | want: %v", err, test.error)
			}
		})
	}
}
