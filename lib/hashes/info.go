package hashes

import (
	"hash"
    "golang.org/x/crypto/sha3"
)

var (
	hashFunction hash.Hash = sha3.New256()
)

type InfoHash struct{
	data []byte
	}

func getIdentifier(data []byte) InfoHash {
	return InfoHash{hashFunction.Sum(data)}
}

// Xor is a function that xors two InfoHashes. This is needed for Kademlias XOR
// metric. The result is a new InfoHash
//func (i InfoHash) Xor(another InfoHash) InfoHash {
//	var hash InfoHash
//	for b := range i {
//		hash[b] = i[b] ^ another[b]
//	}
//	return hash
//}

// CommonPrefixLenght compares to InfoHashes bitwise and finds the length of the
// bit-prefix both have in common.
//func CommonPrefixLength(a, b InfoHash) int {
//	var length, sum int
//	for length = 0; length < len(a); length++ {
//		if a[length] != b[length] {
//			bits := a[length] | b[length]
//			var i uint = 1
//			for i <= 8 {
//				if (bits >> i) == 0 {
//					sum += int(8 - i)
//					break
//				}
//				i++
//			}
//			break
//		} else {
//			sum += 8
//		}
//	}
//	return sum
//}

// Implementation of the GoStringer interface
//func (i InfoHash) GoString() string {
//	return fmt.Sprintf("%#v", i[:])
//}

// Implementation of the Stringer interface. All hashes are written as their
// base64-encoded content
//func (i InfoHash) String() string {
//	return base64.StdEncoding.EncodeToString(i[:])
//}
