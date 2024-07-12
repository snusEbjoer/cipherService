package encodeutils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

type Hash struct {
	hash []byte
}

func (h Hash) ToString() string {
	return hex.EncodeToString(h.hash)
}

func (h Hash) ToHash() []byte {
	return h.hash
}

func NewSHA256(data string) Hash {
	hash := sha256.Sum256([]byte(data))
	return Hash{
		hash: hash[:],
	}
}

func NewMD5(data string) Hash {
	hash := md5.Sum([]byte(data))
	return Hash{
		hash: hash[:],
	}
}
