package marker

import (
	"fmt"
	"github.com/0chain/gosdk/core/encryption"
	"github.com/0chain/gosdk/zboxcore/client"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)


func TestWriteMarker_GetHashData(t *testing.T) {
	allocRoot := "528c04e3143f0cf14384fe0d04dd8af3e7764a2344886bebbbd3a3b1641cee26"
	prevAllocRoot := "9bf430d6f086f1bdc2d26ad7a708a0e7958aa9ae20efbc6778450739fb1ca468"
	allocID := "7827c760363c4836b0acd3967a023c9061deef50e0c44db3e3aca8e14f8ef6f8"
	blobberID := "4a0ffbd42c64f44ec1cca858c7e5b5fd408911ed03df3b7009049cdb76e03ac2"
	clientID := "b6de562b57a0b593d0480624f79a55ed46dba544404595bee0273144e01034ae"
	size := rand.Int63()
	timestamp := time.Now()

	wm := &WriteMarker{
		AllocationRoot:         allocRoot,
		PreviousAllocationRoot: prevAllocRoot,
		AllocationID:           allocID,
		Size:                   size,
		BlobberID:              blobberID,
		Timestamp:              timestamp.Unix(),
		ClientID:               clientID,
		Signature:              "",
	}
	expectedData := fmt.Sprintf("%v:%v:%v:%v:%v:%v:%v", allocRoot, prevAllocRoot, allocID, blobberID, clientID, size, timestamp.Unix())
	assert.Equal(t, expectedData, wm.GetHashData())
}

func TestWriteMarker_GetHash(t *testing.T) {
	allocRoot := "528c04e3143f0cf14384fe0d04dd8af3e7764a2344886bebbbd3a3b1641cee26"
	prevAllocRoot := "9bf430d6f086f1bdc2d26ad7a708a0e7958aa9ae20efbc6778450739fb1ca468"
	allocID := "7827c760363c4836b0acd3967a023c9061deef50e0c44db3e3aca8e14f8ef6f8"
	blobberID := "4a0ffbd42c64f44ec1cca858c7e5b5fd408911ed03df3b7009049cdb76e03ac2"
	clientID := "b6de562b57a0b593d0480624f79a55ed46dba544404595bee0273144e01034ae"
	size := rand.Int63()
	timestamp := time.Now()

	wm := &WriteMarker{
		AllocationRoot:         allocRoot,
		PreviousAllocationRoot: prevAllocRoot,
		AllocationID:           allocID,
		Size:                   size,
		BlobberID:              blobberID,
		Timestamp:              timestamp.Unix(),
		ClientID:               clientID,
		Signature:              "",
	}
	expectedHash := encryption.Hash(fmt.Sprintf("%v:%v:%v:%v:%v:%v:%v", allocRoot, prevAllocRoot, allocID, blobberID, clientID, size, timestamp.Unix()))
	assert.Equal(t, expectedHash, wm.GetHash())
}

func TestWriteMarker_Sign(t *testing.T) {
	walletConf := `{"client_id":"9bf430d6f086f1bdc2d26ad7a708a0e7958aa9ae20efbc6778450739fb1ca468","client_key":"eeb0c33325cbee0fb58bc09962f69a44d0b22ac2824a063eb1002273347e601a4612e6fea7e1a1ae62e0e3b7f1301c4de8a855bae86ebfa6e9dbbb41c3e39c24","keys":[{"public_key":"eeb0c33325cbee0fb58bc09962f69a44d0b22ac2824a063eb1002273347e601a4612e6fea7e1a1ae62e0e3b7f1301c4de8a855bae86ebfa6e9dbbb41c3e39c24","private_key":"0800424da684ff94ac8af3ccc3e024a8d16bb9054237e79feffc486efda6e210"}],"mnemonics":"neck hurt glow action goose gadget meat ankle patch boy truth convince glass grief cloth sunny evil puppy decorate language okay burst replace cigar","version":"1.0","date_created":"2021-03-13 01:42:49.625529496 +0700 +07 m=+1.336039112"}`
	err := client.PopulateClient(walletConf, "bls0chain")
	assert.NoErrorf(t, err, "unexpected client.PopulateClient() error but got %v", err)
	allocRoot := "528c04e3143f0cf14384fe0d04dd8af3e7764a2344886bebbbd3a3b1641cee26"
	prevAllocRoot := "9bf430d6f086f1bdc2d26ad7a708a0e7958aa9ae20efbc6778450739fb1ca468"
	allocID := "7827c760363c4836b0acd3967a023c9061deef50e0c44db3e3aca8e14f8ef6f8"
	blobberID := "4a0ffbd42c64f44ec1cca858c7e5b5fd408911ed03df3b7009049cdb76e03ac2"
	clientID := "b6de562b57a0b593d0480624f79a55ed46dba544404595bee0273144e01034ae"
	size := rand.Int63()
	timestamp := time.Now()

	wm := &WriteMarker{
		AllocationRoot:         allocRoot,
		PreviousAllocationRoot: prevAllocRoot,
		AllocationID:           allocID,
		Size:                   size,
		BlobberID:              blobberID,
		Timestamp:              timestamp.Unix(),
		ClientID:               clientID,
		Signature:              "",
	}

	expectedSignature, err := client.Sign(encryption.Hash(fmt.Sprintf("%v:%v:%v:%v:%v:%v:%v", allocRoot, prevAllocRoot, allocID, blobberID, clientID, size, timestamp.Unix())))
	assert.NoErrorf(t, err, "unexpected error but got %v", err)
	err = wm.Sign()
	assert.NoErrorf(t, err, "unexpected rm.Sign() error but got %v", err)
	assert.Equal(t, expectedSignature, wm.Signature)
}