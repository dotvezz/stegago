package steganographer

import (
	"image"
	"os"
	"stegago/engines"
)

// New returns a Steganographer with defaults: nil Processor and LSBContainer Engine
func New() Steganographer {
	return Steganographer{
		e: &engines.LSBContainer{},
	}
}

// Open returns a new Stenographer with defaults and an image loaded from the path. An error
// is returned if unable to open the image
func Open(path string) (s Steganographer, err error) {
	f, err := os.Open(path)

	s := New()

	return
}

// Load returns a new stenographer with the defaults and an image loaded directly from the
// image.Image passed as the only parameter
func Load(i image.Image) (s Steganographer) {
	s = New()
	s.i = &i
	return
}

// Stenographer is the primary means of using the stegago package. It holds internal references
type Steganographer struct {
	e Engine
	p Processor
}

// Load
func (s *Steganographer) Load(i image.Image) {
	s.i = &i
}
