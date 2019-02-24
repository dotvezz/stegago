package lsb

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
)

type Engine struct{}

func (Engine) Encode(im *image.Image, in []byte) (err error) {
	if im == nil {
		return errors.New("image must not be nil")
	}
	rgba, ok := (*im).(*image.RGBA)
	if !ok {
		toRGBA(im)
		rgba, ok = (*im).(*image.RGBA)
		if !ok {
			return errors.New("unable to copy image to RGBA in Engine Engine")
		}
	}

	w := rgba.Bounds().Dx()
	h := rgba.Bounds().Dy()
	max := len(in) * 8

	bits := make([]uint8, max)

	for ind, bte := range in {
		for a := 0; a < 8; a++ {
			bits[ind*8+7-a] = bte & 1
			bte >>= 1
		}
	}

	if max/3 > w*h {
		return errors.New("data is too large to be encoded in this image")
	}

	for z := 0; z < max; z++ {
		y := z / w
		x := z % w
		c := rgba.At(x, y).(color.RGBA)

		b := bits[z]
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
		*ch <<= 1
		*ch += (*ch) & b
		rgba.Set(x, y, c)
	}

	*im = rgba
	return
}

func (Engine) Decode(i *image.Image) (out []byte, err error) {
	//TODO: Actually do something here
	return
}

func bitAt(d *[]byte, i int) uint8 {
	x := i / 8
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
