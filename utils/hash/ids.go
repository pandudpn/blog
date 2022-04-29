package hash

import (
	"os"
	
	"github.com/speps/go-hashids/v2"
)

type hash struct {
	h *hashids.HashID
}

// NewHash is constructor of package hashing
// this is will return an instance of HashID for encode or decode
func NewHash() *hash {
	alphabet := os.Getenv("HASH_IDS_ALPHABET")
	
	hd := hashids.NewData()
	hd.Salt = os.Getenv("SALT_IDS")
	hd.MinLength = 11
	if alphabet != "" {
		hd.Alphabet = alphabet
	}
	
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil
	}
	
	return &hash{
		h: h,
	}
}
