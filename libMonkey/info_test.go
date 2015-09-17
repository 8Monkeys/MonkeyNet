package libMonkey

import "testing"

func TestInfoHashCreation(t *testing.T) {
	var null infoHash
	_, i := newRandom()
	_, j := newRandom()

	if i == null || j == null {
		t.Error("New infoHash objects must not be 0 when initialized by random")
	}
	if i == j {
		t.Error("%v and %v must not be equal when generated with New()")
	}
}
