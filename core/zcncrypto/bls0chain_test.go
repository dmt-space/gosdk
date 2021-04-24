package zcncrypto

import (
	"testing"

	"github.com/0chain/gosdk/core/encryption"
)

import "fmt"
import "github.com/herumi/bls-go-binary/bls"

var verifyPublickey = `e8a6cfa7b3076ae7e04764ffdfe341632a136b52953dfafa6926361dd9a466196faecca6f696774bbd64b938ff765dbc837e8766a5e2d8996745b2b94e1beb9e`
var signPrivatekey = `5e1fc9c03d53a8b9a63030acc2864f0c33dffddb3c276bf2b3c8d739269cc018`
var data = `TEST`
var blsWallet *Wallet

// This is a basic unit test to check that MIRACL generates correct public key.
func TestHerumiPKcompatibility(t *testing.T) {
	var skStr = signPrivatekey
	// var skStr = "57f6332231ed63c0eba947da0054b74367cc19a7c213b651beb7b9f659e703a"
	var sk bls.SecretKey
	sk.DeserializeHexStr(skStr)
	// sk.SetHexString(skStr)
	pk := sk.GetPublicKey()

	skStr2 := sk.GetHexString()
	if skStr2 != skStr {
		// panic("Secret Key deserialize failed: [skStr, skStr2]: " + skStr + " " + skStr2)
	}

	fmt.Println("sk", skStr2)
	fmt.Println("pk", pk.SerializeToHexStr())
	fmt.Println("pk", pk.GetHexString())
}

// var signature = `032419be0bed64c448a3a5d2e3e98e295da99f6b442c657f10ae2514cc050e259d`
var signature = `031fac7be9d577f3c813c27c90b1a33f0a6286399d01bb89bcd3a0ba1b3f1691ca`
func TestSignatureDeserialize(t *testing.T) {
	var sig bls.Sign
	err := sig.DeserializeHexStr(signature)
	if err != nil {
		panic("ruh roh")
	}
}

func TestPubKeySerialize3(t *testing.T) {
	var pk bls.PublicKey
	bls.BlsGetGeneratorOfPublicKey(&pk)
	fmt.Println("base?", pk.GetHexString())
}

var miracl2PK = `1 025bdd03aca35fd8816993afc6e8d1e8334341c6b8d0859ea680ca983c2c4f4e 1bc4bbda762b218d2a7d2d9450476086d6d9c4d3b2341070dbed74d63e5bd21f 03ae589093de94f88f4a3544990d81a2e6b65f0d86ac746559ded9d06653337e 060a5ce375bc7789f0ccfb778ba206121a0d755137f51defc3ae3c59addabf29`
var herumiPK = `1 1bdfed3a85690775ee35c61678957aaba7b1a1899438829f1dc94248d87ed368 18a02c6bd223ae0dfda1d2f9a3c81726ab436ce5e9d17c531ff0a385a13a0b49 39ac7dfc3364e851ebd2631ea6f1685609fc66d50223cc696cb59ff2fee47ac 17f6dfafec19bfa87bf791a4d694f43fec227ae6f5a867490e30328cac05eaff`

var miraclPK = `1 03ae589093de94f88f4a3544990d81a2e6b65f0d86ac746559ded9d06653337e 060a5ce375bc7789f0ccfb778ba206121a0d755137f51defc3ae3c59addabf29 025bdd03aca35fd8816993afc6e8d1e8334341c6b8d0859ea680ca983c2c4f4e 1bc4bbda762b218d2a7d2d9450476086d6d9c4d3b2341070dbed74d63e5bd21f`
func TestPubKeySerialize4(t *testing.T) {
	var pk bls.PublicKey
	pk.SetHexString(miraclPK)
	fmt.Println("pk", pk.GetHexString())
}

func TestPubKeySerialize2(t *testing.T) {
	var pk bls.PublicKey

	pk.SetHexString(miraclPK)
	fmt.Println("pk", pk.GetHexString())

	pk.SetHexString(miracl2PK)
	fmt.Println("pk", pk.GetHexString())

	pk.SetHexString(herumiPK)
	fmt.Println("pk", pk.GetHexString())

	bls.BlsGetGeneratorOfPublicKey(&pk)
	fmt.Println("base?", pk.GetHexString())
	// pub2 := new(bls.PublicKey)
	// C.blsGetGeneratorOfG2(pub2.getPointer())
	// fmt.Println("base?", pub2.GetHexString())
}

// Basic unit test for de/serialization of herumi public keys.
func TestPubKeySerialize(t *testing.T) {
	var sk bls.SecretKey
	sk.SetHexString("057f6332231ed63c0eba947da0054b74367cc19a7c213b651beb7b9f659e703a")
	pk := sk.GetPublicKey()
	fmt.Println("sk", sk.GetHexString())
	fmt.Println("pk", pk.GetHexString())
}

	// // fmt.Println("random7"); // Changing this string forces a new random SK.
	// sk.SetByCSPRNG()

func TestSignatureScheme(t *testing.T) {
	sigScheme := NewSignatureScheme("bls0chain")
	switch sigScheme.(type) {
	case SignatureScheme:
		// pass
	default:
		t.Fatalf("Signature scheme invalid")
	}
	w, err := sigScheme.GenerateKeys()
	if err != nil {
		t.Fatalf("Generate Key failed %s", err.Error())
	}
	if w.ClientID == "" || w.ClientKey == "" || len(w.Keys) != 1 || w.Mnemonic == "" {
		t.Fatalf("Invalid keys generated")
	}
	blsWallet = w
}

func TestSSSignAndVerify(t *testing.T) {
	fmt.Println("v5")
	signScheme := NewSignatureScheme("bls0chain")
	signScheme.SetPrivateKey(signPrivatekey)
	hash := Sha3Sum256(data)
	fmt.Println("hash to compare:", hash)
	signature, err := signScheme.Sign(hash)

	var sk bls.SecretKey
	sk.DeserializeHexStr(signPrivatekey)
	fmt.Println("secretkey to use for miracl", sk.GetHexString())

	var sig bls.Sign
	sig.DeserializeHexStr(signature)
	fmt.Println("signature to compare:", signature)
	fmt.Println("signature to compare:", sig.GetHexString())

	if err != nil {
		t.Fatalf("BLS signing failed")
	}
	verifyScheme := NewSignatureScheme("bls0chain")
	verifyScheme.SetPublicKey(verifyPublickey)
	if ok, err := verifyScheme.Verify(signature, hash); err != nil || !ok {
		t.Fatalf("Verification failed\n")
	}
}

func BenchmarkBLSSign(b *testing.B) {
	sigScheme := NewSignatureScheme("bls0chain")
	sigScheme.SetPrivateKey(signPrivatekey)
	for i := 0; i < b.N; i++ {
		_, err := sigScheme.Sign(encryption.Hash(data))
		if err != nil {
			b.Fatalf("BLS signing failed")
		}
	}
}

func TestRecoveryKeys(t *testing.T) {

	sigScheme := NewSignatureScheme("bls0chain")
    TestSignatureScheme(t)
	w, err := sigScheme.RecoverKeys(blsWallet.Mnemonic)
	if err != nil {
		t.Fatalf("set Recover Keys failed")
	}
	if w.ClientID != blsWallet.ClientID || w.ClientKey != blsWallet.ClientKey {
		t.Fatalf("Recover key didn't match with generated keys")
	}

}

func TestCombinedSignAndVerify(t *testing.T) {
	sk0 := `c36f2f92b673cf057a32e8bd0ca88888e7ace40337b737e9c7459fdc4c521918`
	sk1 := `704b6f489583bf1118432fcfb38e63fc2d4b61e524fb196cbd95413f8eb91c12`
	primaryKey := `f72fd53ee85e84157d3106053754594f697e0bfca1f73f91a41f7bb0797d901acefd80fcc2da98aae690af0ee9c795d6590c1808f26490306433b4e9c42f7b1f`

	hash := Sha3Sum256(data)
	// Create signatue for 1
	sig0 := NewSignatureScheme("bls0chain")
	err := sig0.SetPrivateKey(sk0)
	if err != nil {
		t.Fatalf("Set private key failed - %s", err.Error())
	}
	signature, err := sig0.Sign(hash)
	if err != nil {
		t.Fatalf("BLS signing failed")
	}
	// Create signature for second
	sig1 := NewSignatureScheme("bls0chain")
	err = sig1.SetPrivateKey(sk1)
	if err != nil {
		t.Fatalf("Set private key failed - %s", err.Error())
	}
	addSig, err := sig1.Add(signature, hash)

	verifyScheme := NewSignatureScheme("bls0chain")
	err = verifyScheme.SetPublicKey(primaryKey)
	if err != nil {
		t.Fatalf("Set public key failed")
	}
	if ok, err := verifyScheme.Verify(addSig, hash); err != nil || !ok {
		t.Fatalf("Verification failed\n")
	}
}

func TestSplitKey(t *testing.T) {
	primaryKeyStr := `c36f2f92b673cf057a32e8bd0ca88888e7ace40337b737e9c7459fdc4c521918`
	sig0 := NewBLS0ChainScheme()
	err := sig0.SetPrivateKey(primaryKeyStr)
	if err != nil {
		t.Fatalf("Set private key failed - %s", err.Error())
	}
	hash := Sha3Sum256(data)
	signature, err := sig0.Sign(hash)
	if err != nil {
		t.Fatalf("BLS signing failed")
	}
	numSplitKeys := int(2)
	w, err := sig0.SplitKeys(numSplitKeys)
	if err != nil {
		t.Fatalf("Splitkeys key failed - %s", err.Error())
	}
	sigAggScheme := make([]BLS0ChainScheme, numSplitKeys)
	for i := 0; i < numSplitKeys; i++ {
		sigAggScheme[i].SetPrivateKey(w.Keys[i].PrivateKey)
	}
	var aggrSig string
	for i := 1; i < numSplitKeys; i++ {
		tmpSig, _ := sigAggScheme[i].Sign(hash)
		aggrSig, _ = sigAggScheme[0].Add(tmpSig, hash)
	}
	if aggrSig != signature {
		t.Fatalf("split key signature failed")
	}
}
