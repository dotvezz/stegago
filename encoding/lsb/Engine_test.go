package lsb

import (
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"os"
	"testing"
)

func BenchmarkLsbContainer_Encode(b *testing.B) {
	i, err := loadImage()
	if err != nil {
		b.Error(err)
	}

	engine := Engine{}
	for x := 0; x < b.N; x++ {
		err = engine.Encode(i, []byte("Hello, world!"))
		if err != nil {
			b.Error(err)
		}
	}

	saveImage(*i)
}

func loadImage() (i *image.Image, err error) {
	f, err := os.Open("../../testImages/sailboat.jpg")
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

func saveImage(i image.Image) {
	w, _ := os.Create("../../testImages/sailboat.out.jpg")
	jpeg.Encode(w, i, nil)
}
