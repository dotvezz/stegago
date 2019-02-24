package endocing

import (
	"image"
	"stegago"
)

type Decoder func(d *[]byte, i *image.Image) (err error)

func NewDecoder(engine stegago.Engine) Decoder {
	return func(d *[]byte, i *image.Image) (err error) {

	}
}
