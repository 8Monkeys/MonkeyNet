package libMonkey

import "testing"

func TestInfoHashFromRandom(t *testing.T) {
	i, e := NewRandom()
	if e != nil {
		t.Error("Failed to create InfoHash: %v", e)
	}

	j, e := NewRandom()

	if i.Empty() || j.Empty() {
		t.Error("New InfoHash objects must not be Empty when initialized by random")
	}
	if i == j {
		t.Error("%v and %v must not be equal when generated with New()", i, j)
	}
}

func TestInfoDefaultInit(t *testing.T) {
	var i InfoHash
	if !i.Empty() {
		t.Error("Default initialisation failed. Should be '0', was %s", i)
	}
}

func TestIoWriterImplementation(t *testing.T) {
	buffers := [][]byte{
		[]byte(""),
		[]byte("1234567890"),
		[]byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")}

	for b := range buffers {
		var i InfoHash
		_, e := i.Write(buffers[b])
		if e != nil {
			t.Error("Error during writing to InfoHash:", e)
		}
	}
}

func TestWritingToInitializedInfoHash(t *testing.T) {
	i, _ := NewRandom()
	_, e := i.Write([]byte("a dummy"))
	if e == nil {
		t.Error("Failed to warn on writing to already initialised InfoHash:", e, "i was ", i)
	}
}
