package libMonkey

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

type InfoHash [32]byte

func NewRandom() (InfoHash, error) {
	var newling InfoHash
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return newling, err
	}

	for i := range b {
		newling[i] = b[i]
	}
	return newling, nil
}

func (i *InfoHash) Empty() bool {
	for it := range i {
		if i[it] != 0 {
			return false
		}
	}
	return true
}

func (i *InfoHash) Write(p []byte) (n int, err error) {
	if !i.Empty() {
		return 0, fmt.Errorf("%s is an already initialized value", i)
	}

	err = nil
	var it, ti int
	it = len(i)
	if len(i) < len(p) {
		n = len(i)
		ti = len(i)
	} else {
		n = len(p)
		ti = len(p)
	}
	for it > 0 && ti > 0 {
		i[it-1] = p[ti-1]
		it--
		ti--
	}

	return
}

func (i InfoHash) Xor(another InfoHash) InfoHash {
	var hash InfoHash
	for b := range i {
		hash[b] = i[b] ^ another[b]
	}
	return hash
}

func (i InfoHash) GoString() string {
	return base64.StdEncoding.EncodeToString(i[:])
}

func (i InfoHash) String() string {
	return i.GoString()
}
