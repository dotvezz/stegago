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
	rgba, ok := (*i).(*image.RGBA)
	if !ok {
		toRGBA(i)
		rgba, ok = (*i).(*image.RGBA)
		if !ok {
			return errors.New("unable to copy image to RGBA in LSB Engine")
		}
	}

	w := rgba.Bounds().Dx()
	h := rgba.Bounds().Dy()

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := rgba.At(x, y).(color.RGBA)
			z := w*x + y
			b := bitAt(d, z)
			if b > 1 {
				goto Fin
			}
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
			*ch <<= b
			rgba.Set(x, y, c)
		}
	}

	Fin:
	*i = rgba
	return
}

func (lsbContainer) Decode(d *[]byte, i *image.Image) (err error) {
	//TODO: Actually do something here
	return
}

func bitAt(d *[]byte, i int) uint8 {
	x := i/8
	if len(*d) <= x {
		return 2
	}
	b := (*d)[x]
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
