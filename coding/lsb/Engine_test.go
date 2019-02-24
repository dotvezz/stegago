package lsb

import (
	"image"
	_ "image/jpeg"
	"os"
	"testing"
)

func BenchmarkLsbContainer_Encode(b *testing.B) {
	i, err := loadImage()
	if err != nil {
		b.Error(err)
	}
	var message []byte
	message = []byte("Hello, world!")
	engine := lsbContainer{}
	for x := 0; x < b.N; x++ {
		err = engine.Encode(i, &message)
		if err != nil {
			b.Error(err)
		}
	}
}

func loadImage() (i *image.Image, err error) {
	f, err := os.Open("../images/sailboat.jpg")
	if err != nil {
		return
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return
	}
	i = &img
	return
}