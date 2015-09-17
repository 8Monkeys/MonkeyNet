package libMonkey

import "crypto/rand"

type infoHash [32]byte

func newRandom() (error, infoHash) {
	var newling infoHash
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return err, newling
	}

	for i := range b {
		newling[i] = b[i]
		}
	return nil, newling
}
