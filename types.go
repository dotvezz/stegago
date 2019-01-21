package stegago

import "image"

type Processor interface {
	Post(d *[]byte) (err error)
	Pre(d *[]byte) (err error)
}

type Engine interface {
	Encode(image *image.Image, d *[]byte) (err error)
	Decode(image *image.Image, d *[]byte) (err error)
}
