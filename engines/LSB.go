package engines

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
)

var LSB = &lsbContainer{}

type lsbContainer struct{}

func (lsbContainer) Encode(i *image.Image, d *[]byte) (err error) {
	toRGBA(i)
	rgba, ok := (*i).(*image.RGBA)
	if !ok {
		return errors.New("unable to copy image to RGBA in LSB Engine")
	}

	w := rgba.Bounds().Dx()
	h := rgba.Bounds().Dy()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := rgba.At(x, y).(color.RGBA)
			z := w*x + y
			var ch *uint8
			switch z % 3 {
			case 0:
				ch = &c.R
				break
			case 1:
				ch = &c.G
				break
			case 2:
				ch = &c.B
				break
			}
			*ch >>= 1
			*ch <<= bitAt(d, z)
			rgba.Set(x, y, c)
		}
	}

	*i = rgba
	return
}

func (lsbContainer) Decode(d *[]byte, i *image.Image) (err error) {
	//TODO: Actually do something here
	return
}

func bitAt(d *[]byte, i int) uint8 {
	b := (*d)[i/8]
	b >>= uint8(7 - i%8)
	return 1 & b
}

func toRGBA(i *image.Image) {
	in := *i
	b := in.Bounds()
	out := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(out, out.Bounds(), in, b.Min, draw.Src)
	*i = out
}
