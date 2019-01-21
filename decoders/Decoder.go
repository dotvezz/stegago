package decoders

import (
	"image"
	"stegago"
)

type Decoder func(i *image.Image, d *[]byte) (err error)

func New(engine stegago.Engine) Decoder {
	return func(i *image.Image, d *[]byte) (err error) {

	}
}
