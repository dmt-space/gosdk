// +build !js,!wasm

package zcncrypto

import (
	"encoding/hex"
	"fmt"
	"testing"

	BN254 "github.com/0chain/gosdk/miracl"
	"github.com/herumi/bls-go-binary/bls"
	"github.com/stretchr/testify/require"
)

func TestHerumiToMiraclPK(t *testing.T) {

}

func TestMiraclToHerumiPK(t *testing.T) {
	miraclpk1 := `0418a02c6bd223ae0dfda1d2f9a3c81726ab436ce5e9d17c531ff0a385a13a0b491bdfed3a85690775ee35c61678957aaba7b1a1899438829f1dc94248d87ed36817f6dfafec19bfa87bf791a4d694f43fec227ae6f5a867490e30328cac05eaff039ac7dfc3364e851ebd2631ea6f1685609fc66d50223cc696cb59ff2fee47ac`
	pk1 := MiraclToHerumiPK(miraclpk1)

	require.EqualValues(t, pk1, "68d37ed84842c91d9f82389489a1b1a7ab7a957816c635ee750769853aeddf1b490b3aa185a3f01f537cd1e9e56c43ab2617c8a3f9d2a1fd0dae23d26b2ca018")

	// Assert DeserializeHexStr works on the output of MiraclToHerumiPK
	var pk bls.PublicKey
	err := pk.DeserializeHexStr(pk1)
	require.NoError(t, err)

	fmt.Println(pk.GetHexString())
}

func TestDecodeHerumiPK(t *testing.T) {
	rawPK := "1 1bdfed3a85690775ee35c61678957aaba7b1a1899438829f1dc94248d87ed368 18a02c6bd223ae0dfda1d2f9a3c81726ab436ce5e9d17c531ff0a385a13a0b49 39ac7dfc3364e851ebd2631ea6f1685609fc66d50223cc696cb59ff2fee47ac 17f6dfafec19bfa87bf791a4d694f43fec227ae6f5a867490e30328cac05eaff"
	hexPK := "e8a6cfa7b3076ae7e04764ffdfe341632a136b52953dfafa6926361dd9a466196faecca6f696774bbd64b938ff765dbc837e8766a5e2d8996745b2b94e1beb9e"

	b, err := hex.DecodeString(hexPK)

	b2 := BN254.FromBytes(b)
	p := b2.ToString()
	fmt.Println(p)
	pk, err := DeserializeHerumiPK(hexPK)

	require.NoError(t, err)
	require.EqualValues(t, rawPK, pk)

}

var herumiPublicKey = "1 1966a4d91d362669fafa3d95526b132a6341e3dfff6447e0e76a07b3a7cfa6e8 1eeb1b4eb9b2456799d8e2a566877e83bc5d76ff38b964bd4b7796f6a6ccae6f 46d6e633f5eb68a93013dfac1420bf7a1e1bf7a87476024478e97a1cc115de9 34574266b382b8e5174477ab8a32a49a57eda74895578031cd2d41fd0aef446"
var miraclPublicKey = "1966a4d91d362669fafa3d95526b132a6341e3dfff6447e0e76a07b3a7cfa6e81eeb1b4eb9b2456799d8e2a566877e83bc5d76ff38b964bd4b7796f6a6ccae6f46d6e633f5eb68a93013dfac1420bf7a1e1bf7a87476024478e97a1cc115de934574266b382b8e5174477ab8a32a49a57eda74895578031cd2d41fd0aef446"
