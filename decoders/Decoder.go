package decoders

import (
	"image"
	"stegago"
)

type Decoder func(d *[]byte, i image.Image) (err error)

func New(engine stegago.Engine) Decoder {
	return func(i *image.Image, d *[]byte) (err error) {

	}
}
