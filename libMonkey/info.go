package libMonkey

import (
	"crypto/rand"
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
	fmt.Printf("Trying to write '%s' to '%s'\n", p, i)

	if !i.Empty() {
		fmt.Printf("<Failed>\n\n")
		return 0, fmt.Errorf("%s is an already initialized value", i)
	}

	it := 0
	for it < len(p) && it < len(i) {
		i[it] = p[it]
		it++
	}

	fmt.Printf("i is now '%s', remainder of p is '%s'\n\n", i, p[it:])

	return it, nil
}
