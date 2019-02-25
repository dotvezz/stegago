package lsb

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
)

type Engine [0]bool

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
	in = append(in, byte(0))
	if len(in)*8 > w*h {
		return errors.New("data is too large to be encoded in this image")
	}

	z := 0
	for _, bte := range in {
		for a := 0; a < 8; a++ {
			b := bte & 1
			bte >>= 1

			y := z / 3 / w
			x := z / 3 % w
			c := rgba.At(x, y).(color.RGBA)

			var ch *uint8
			switch z % 3 {
			case 0:
				ch = &c.R
			case 1:
				ch = &c.G
			case 2:
				ch = &c.B
			}

			*ch >>= 1
			*ch <<= 1
			*ch += (*ch) & b
			rgba.Set(x, y, c)

			z++
		}
	}

	*im = rgba
	return
}

func (Engine) Decode(im *image.Image) (out []byte, err error) {
	if im == nil {
		return nil, errors.New("image must not be nil")
	}

	rgba, ok := (*im).(*image.RGBA)
	if !ok {
		toRGBA(im)
		rgba, ok = (*im).(*image.RGBA)
		if !ok {
			return nil, errors.New("unable to copy image to RGBA in Engine Engine")
		}
	}

	byt := byte(0)

	for i, p := range rgba.Pix {
		byt <<= 1
		byt = byt | (p & 1)
		if i > 0 && i % 8 == 0 {
			if byt == byte(0) {
				break
			}
			out = append(out, byt)
			byt = byte(0)
		}
	}

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
