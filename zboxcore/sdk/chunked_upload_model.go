package sdk

import (
	"hash/fnv"
	"strconv"

	"github.com/0chain/gosdk/core/util"
	"github.com/0chain/gosdk/zboxcore/fileref"
)

// FileMeta metadata of stream input/local
type FileMeta struct {
	// Mimetype mime type of source file
	MimeType string

	// Path local path of source file
	Path string
	// ThumbnailPath local path of source thumbnail
	ThumbnailPath string

	// ActualHash hash of orignial file (unencoded, unencrypted)
	ActualHash string
	// ActualSize total bytes of  orignial file (unencoded, unencrypted).  it is 0 if input is live stream.
	ActualSize int64
	// ActualThumbnailSize total bytes of orignial thumbnail (unencoded, unencrypted)
	ActualThumbnailSize int64
	// ActualThumbnailHash hash of orignial thumbnail (unencoded, unencrypted)
	ActualThumbnailHash string

	//RemoteName remote file name
	RemoteName string
	// RemotePath remote path
	RemotePath string
	// Attributes file attributes in blockchain
	Attributes fileref.Attributes
}

// FileID generante id of progress on local cache
func (meta *FileMeta) FileID() string {

	hash := fnv.New64a()
	hash.Write([]byte(meta.Path + "_" + meta.RemotePath))

	return strconv.FormatUint(hash.Sum64(), 36) + "_" + meta.RemoteName
}

// UploadFormData form data of upload
type UploadFormData struct {
	ConnectionID string `json:"connection_id,omitempty"`
	// Filename remote file name
	Filename string `json:"filename,omitempty"`
	// Path remote path
	Path string `json:"filepath,omitempty"`

	// ContentHash hash of shard data (encoded,encrypted) when it is last chunk. it is ChunkHash if it is not last chunk.
	ContentHash string `json:"content_hash,omitempty"`
	// Hash hash of shard thumbnail  (encoded,encrypted)
	ThumbnailContentHash string `json:"thumbnail_content_hash,omitempty"`

	// MerkleRoot challenge hash of shard data (encoded, encrypted)
	MerkleRoot string `json:"merkle_root,omitempty"`

	// ActualHash hash of orignial file (unencoded, unencrypted)
	ActualHash string `json:"actual_hash,omitempty"`
	// ActualSize total bytes of  orignial file (unencoded, unencrypted)
	ActualSize int64 `json:"actual_size,omitempty"`
	// ActualThumbnailSize total bytes of orignial thumbnail (unencoded, unencrypted)
	ActualThumbSize int64 `json:"actual_thumb_size,omitempty"`
	// ActualThumbnailHash hash of orignial thumbnail (unencoded, unencrypted)
	ActualThumbHash string `json:"actual_thumb_hash,omitempty"`

	MimeType     string             `json:"mimetype,omitempty"`
	CustomMeta   string             `json:"custom_meta,omitempty"`
	EncryptedKey string             `json:"encrypted_key,omitempty"`
	Attributes   fileref.Attributes `json:"attributes,omitempty"`

	IsFinal      bool   `json:"is_final,omitempty"`      // current chunk is last or not
	ChunkHash    string `json:"chunk_hash"`              // hash of current chunk
	ChunkIndex   int    `json:"chunk_index,omitempty"`   // the seq of current chunk. all chunks MUST be uploaded one by one because of streaming merkle hash
	ChunkSize    int64  `json:"chunk_size,omitempty"`    // the size of a chunk. 64*1024 is default
	UploadOffset int64  `json:"upload_offset,omitempty"` // It is next position that new incoming chunk should be append to

}

// UploadProgress progress of upload
type UploadProgress struct {
	ID string `json:"id"`

	// ChunkSize size of chunk
	ChunkSize int `json:"chunk_size,omitempty"`
	// EncryptOnUpload encrypt data on upload or not
	EncryptOnUpload  bool `json:"is_encrypted,omitempty"`
	EncryptPrivteKey string

	// ConnectionID chunked upload connection_id
	ConnectionID string `json:"connection_id,omitempty"`
	// ChunkIndex index of last updated chunk
	ChunkIndex int `json:"chunk_index,omitempty"`
	// UploadLength total bytes that has been uploaded to blobbers
	UploadLength int64 `json:"-"`

	Blobbers []*UploadBlobberStatus `json:"merkle_hashers,omitempty"`
}

// UploadBlobberStatus the status of blobber's upload progress
type UploadBlobberStatus struct {
	ChallengeHasher *util.FixedMerkleTree   `json:"challenge_hasher"`
	ContentHasher   *util.CompactMerkleTree `json:"content_hashser"`

	// UploadLength total bytes that has been uploaded to blobbers
	UploadLength int64 `json:"upload_length,omitempty"`
}

// getChallengeHash see detail on https://github.com/0chain/blobber/wiki/Protocols
func (status *UploadBlobberStatus) getChallengeHash() string {
	if status != nil && status.ChallengeHasher != nil {
		return status.ChallengeHasher.GetMerkleRoot()
	}

	return ""

}

// getContentHash see detail on https://github.com/0chain/blobber/wiki/Protocols
func (status *UploadBlobberStatus) getContentHash() string {
	if status != nil && status.ContentHasher != nil {
		return status.ContentHasher.GetMerkleRoot()
	}

	return ""
}