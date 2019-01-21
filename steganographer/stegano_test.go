package steganographer

import "testing"

var sailboatPath = "../images/sailboat.jpg"

func testOpen(t *testing.T) {
	s, err := Open(sailboatPath)
	if err != nil {
		t.Errorf("Unable to open image: %s", sailboatPath)
	}
}
