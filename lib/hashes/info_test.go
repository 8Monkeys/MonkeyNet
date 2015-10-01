package hashes

import (
	"testing"
)

func TestHashing(t *testing.T) {
	info := getIdentifier([]byte("Testing String"))
	if info.data == nil {
		t.Error("The Hash of something can't be nothing")
		}
	t.Logf("info of 'Testing String' is %v", info.data)
}

//func TestInfoHashXOR(t *testing.T) {
//	var hash, empty, another InfoHash
//	if !(empty.Xor(hash) == another) {
//		t.Error("empty XOR empty must always equal empty")
//	}
//	hash, _ = NewRandom()
//	if !(empty.Xor(hash) == hash) {
//		t.Error("empty XOR hash must always equal hash")
//	}
//	if !(hash.Xor(hash) == empty) {
//		t.Error("hash XOR hash must always equal hash")
//	}
//	another, _ = NewRandom()
//	xor := hash.Xor(another)
//	if xor == hash && xor == another {
//		t.Error("hash XOR another must not equal one of them")
//	}
//}

//func TestInfoHashCommonPrefix(t *testing.T) {
//	var hash, empty InfoHash
//	if !(CommonPrefixLength(empty, empty) == 32*8) {
//		t.Error("Common prefix of empty hashes is 32*8 bits long")
//	}
//	hash.Write([]byte("1")) // 1 == bx00110001
//	if CommonPrefixLength(hash, empty) != (31*8)+2 {
//		t.Errorf("%v\n%v\nCommon Prefix != 250", hash, empty)
//	}
//	empty.Write([]byte("O2")) // 4 == bx01001111, 2 == bx00110010
//	if CommonPrefixLength(hash, empty) != (30*8)+1 {
//		t.Errorf("%v\n%v\nCommon Prefix !=242", hash, empty)
//	}
//}
