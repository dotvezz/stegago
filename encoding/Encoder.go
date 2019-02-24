package encoding

import (
	"image"
	"stegago"
)

type Encoder func(i *image.Image, d *[]byte) (err error)

func NewEncoder(engine stegago.Engine) Encoder {
	return func(i *image.Image, d *[]byte) (err error) {
		return engine.Encode(i, d)
	}
}

func (e Encoder) WithProcessor(p stegago.Processor) Encoder {
	return func(i *image.Image, d *[]byte) (err error) {
		err = p.Pre(d)
		if err != nil {
			return
		}
		return e(i, d)
	}
}
