package md5

import (
	"testing"
)

func TestMD5(t *testing.T) {
	data := []byte("Hello there!")
	err := MD5.Pre(&data)

	if err != nil {
		t.Error(err.Error())
	}

	err = MD5.Post(&data)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestMD5Fail(t *testing.T) {
	data := []byte("Hello there!")
	err := MD5.Pre(&data)

	if err != nil {
		t.Error(err.Error())
	}

	// Add a byte so it doesn't match the encoded md5
	data = append(data, 1)

	err = MD5.Post(&data)

	if err == nil {
		t.Error("md5 check passed when it should have failed")
	}
}
