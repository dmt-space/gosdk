package magmasc

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func Test_AccessPoint_Decode(t *testing.T) {
	t.Parallel()

	accessPoint := mockAccessPoint()
	blob, err := json.Marshal(accessPoint)
	if err != nil {
		t.Fatalf("json.Marshal() error: %v | want: %v", err, nil)
	}

	accessPointInvalid := mockAccessPoint()
	accessPointInvalid.Terms = nil
	blobInvalid, err := json.Marshal(accessPointInvalid)
	if err != nil {
		t.Fatalf("json.Marshal() error: %v | want: %v", err, nil)
	}

	tests := [3]struct {
		name  string
		blob  []byte
		want  *AccessPoint
		error bool
	}{
		{
			name:  "OK",
			blob:  blob,
			want:  accessPoint,
			error: false,
		},
		{
			name:  "Decode_ERR",
			blob:  []byte(":"), // invalid json
			want:  &AccessPoint{},
			error: true,
		},
		{
			name:  "Nil_Terms_Invalid_ERR",
			blob:  blobInvalid,
			want:  &AccessPoint{},
			error: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := &AccessPoint{}
			if err := got.Decode(test.blob); (err != nil) != test.error {
				t.Errorf("Decode() error: %v | want: %v", err, test.error)
			}
			if !reflect.DeepEqual(got.Encode(), test.want.Encode()) {
				t.Errorf("Decode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_AccessPoint_Encode(t *testing.T) {
	t.Parallel()

	accessPoint := mockAccessPoint()
	blob, err := json.Marshal(accessPoint)
	if err != nil {
		t.Fatalf("json.Marshal() error: %v | want: %v", err, nil)
	}

	tests := [1]struct {
		name        string
		accessPoint *AccessPoint
		want        []byte
	}{
		{
			name:        "OK",
			accessPoint: accessPoint,
			want:        blob,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.accessPoint.Encode(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Encode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_AccessPoint_Validate(t *testing.T) {
	t.Parallel()

	accessPointNil := mockAccessPoint()
	accessPointNil.AccessPoint = nil

	accessPointTermsNil := mockAccessPoint()
	accessPointTermsNil.Terms = nil

	tests := [3]struct {
		name        string
		accessPoint *AccessPoint
		error       bool
	}{
		{
			name:        "OK",
			accessPoint: mockAccessPoint(),
			error:       false,
		},
		{
			name:        "Access_Point_Is_Nil_ERR",
			accessPoint: accessPointNil,
			error:       true,
		},
		{
			name:        "Terms_Is_Nil_ERR",
			accessPoint: accessPointTermsNil,
			error:       true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if err := test.accessPoint.Validate(); (err != nil) != test.error {
				t.Errorf("Validate() error: %v | want: %v", err, test.error)
			}
		})
	}
}

func Test_AccessPoint_ReadYAML(t *testing.T) {
	t.Parallel()

	var (
		buf = bytes.NewBuffer(nil)
		enc = yaml.NewEncoder(buf)

		accessPoint = mockAccessPoint()
	)

	err := enc.Encode(accessPoint.AccessPoint)
	if err != nil {
		t.Fatalf("yaml Encode() error: %v | want: %v", err, nil)
	}
	path := filepath.Join(t.TempDir(), "config.yaml")
	err = os.WriteFile(path, buf.Bytes(), 0777)
	if err != nil {
		t.Fatalf("os.WriteFile() error: %v | want: %v", err, nil)
	}

	tests := [1]struct {
		name  string
		want  *AccessPoint
		error bool
	}{
		{
			name:  "OK",
			want:  accessPoint,
			error: false,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := &AccessPoint{}
			if err := got.ReadYAML(path); (err != nil) != test.error {
				t.Errorf("ReadYAML() error: %v | want: %v", err, nil)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("ReadYAML() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}
