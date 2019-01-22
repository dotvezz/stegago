package engines

import "image"

var LSB = &lsbContainer{}

type lsbContainer struct{}

func (lsbContainer) Encode(i image.Image, d *[]byte) (err error) {
	if i.
	//TODO: Actually do something here
	return
}

func (lsbContainer) Decode(i image.Image, d *[]byte) (err error) {
	//TODO: Actually do something here
	return
}
