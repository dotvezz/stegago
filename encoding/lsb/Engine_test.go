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
	message := []byte("Hello, world!")
	engine := Engine{}
	for x := 0; x < b.N; x++ {
		err = engine.Encode(i, message)
		if err != nil {
			b.Error(err)
		}
	}

	saveImage(*i)

	i, err = loadSavedImage()
	for x := 0; x < b.N; x++ {
		out, err := engine.Decode(i)
		if err != nil {
			b.Error(err)
		}
		if string(out) != string(message) {
			b.Errorf("%s and %s are different", string(out), string(message))
		}
	}

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

func loadSavedImage() (i *image.Image, err error) {
	f, err := os.Open("../../testImages/sailboat.out.jpg")
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
