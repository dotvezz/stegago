package stegago

import "image"

type Processor interface {
	Post(d *[]byte) (err error)
	Pre(d *[]byte) (err error)
}

type Engine interface {
	Encode(i image.Image, d *[]byte) (err error)
	Decode(i image.Image, d *[]byte) (err error)
}
