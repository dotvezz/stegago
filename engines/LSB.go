package engines

import "image"

var LSB = &lsbContainer{}

type lsbContainer struct{}

func (lsbContainer) Encode(image *image.Image, d *[]byte) (err error) {
	//TODO: Actually do something here
	return
}

func (lsbContainer) Decode(image *image.Image, d *[]byte) (err error) {
	//TODO: Actually do something here
	return
}