package libMonkey

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestInfoHashFromRandom(t *testing.T) {
	i, e := NewRandom()
	if e != nil {
		t.Errorf("Failed to create InfoHash: %v", e)
	}

	j, e := NewRandom()

	if i.Empty() || j.Empty() {
		t.Error("New InfoHash objects must not be Empty when initialized by random")
	}
	if i == j {
		t.Errorf("%v and %v must not be equal when generated with New()", i, j)
	}
}

func BenchmarkInfoHashMathRandGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var iH InfoHash
		for j := 0; j < 31; j++ {
			iH[j] = byte(rand.Intn(256))
		}
	}
}

func BenchmarkInfoHashCurrentRandomGeneration(b *testing.B) {
	var bencher InfoHash
	for i := 0; i < b.N; i++ {
		bencher, _ = NewRandom()
	}
	bencher.Empty()
}

func TestInfoHashDefaultInit(t *testing.T) {
	var i InfoHash
	if !i.Empty() {
		t.Errorf("Default initialisation failed. Should be '0', was %s", i)
	}
}

func TestInfoHashIoWriterImplementation(t *testing.T) {
	buffers := [][]byte{
		[]byte(""),
		[]byte("00000000000000000000000000000000"),
		[]byte("12345678901234567890123456789012"),
		[]byte("1234567890"),
		[]byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")}

	for b := range buffers {
		var i InfoHash
		t.Logf("Writing '% s' to '% s'", buffers[b], i[:])
		_, e := i.Write(buffers[b])
		t.Logf("Got '% s' == %v", i[:], i)
		if e != nil {
			t.Error("Error during writing to InfoHash:", e)
		}
		if len(buffers[b]) != 0 && i.Empty() {
			t.Errorf("InfoHash still empty after writing %v to it.", buffers[b])
		}
	}
}

func TestInfoHashWritingToInitialized(t *testing.T) {
	i, _ := NewRandom()
	_, e := i.Write([]byte("a dummy"))
	if e == nil {
		t.Error("Failed to warn on writing to already initialised InfoHash:", e, "i was ", i)
	}
}

func TestInfoHashFmtWriting(t *testing.T) {
	for i := 0; i < 10; i++ {
		iH, _ := NewRandom()
		str := fmt.Sprintf("%v", iH)
		if str == "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=" {
			t.Errorf("Random InfoHashes must never yield AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=")
		}
	}
	var empty InfoHash
	str := fmt.Sprintf("%v", empty)
	if str != "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=" {
		t.Errorf("Empty InfoHash was %v (should be AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", empty)
	}
}

func TestInfoHashXOR(t *testing.T) {
	var hash, empty, another InfoHash
	if !(empty.Xor(hash) == another) {
		t.Error("empty XOR empty must always equal empty")
	}
	hash, _ = NewRandom()
	if !(empty.Xor(hash) == hash) {
		t.Error("empty XOR hash must always equal hash")
	}
	if !(hash.Xor(hash) == empty) {
		t.Error("hash XOR hash must always equal hash")
	}
	another, _ = NewRandom()
	xor := hash.Xor(another)
	if xor == hash && xor == another {
		t.Error("hash XOR another must not equal one of them")
	}
}

func TestInfoHashCommonPrefix(t *testing.T) {
	var hash, empty InfoHash
	if !(CommonPrefixLength(empty, empty) == 32*8) {
		t.Error("Common prefix of empty hashes is 32*8 bits long")
	}
	hash.Write([]byte("1")) // 1 == bx00110001
	if CommonPrefixLength(hash, empty) != (31*8)+2 {
		t.Errorf("%v\n%v\nCommon Prefix != 250", hash, empty)
	}
	empty.Write([]byte("O2")) // 4 == bx01001111, 2 == bx00110010
	if CommonPrefixLength(hash, empty) != (30*8)+1 {
		t.Errorf("%v\n%v\nCommon Prefix !=242", hash, empty)
	}
}
