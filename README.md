### (Currently WIP)

# Stegago, Steganography in Go

`stegago` is a small package for performing [Steganography][1] in pure Go. It is designed to be
simple to integrate and extensible.

### Basic usage

```Go
package main

import (
	"github.com/dotvezz/img"
	"stegago/encoders"
	"stegago/engines"
	"stegago/processors"
)

func main() {
	hiddenMessage := []byte("Hello, world!")
	image := img.Open("path/to/image.jpg")
	encode := encoders.New(engines.LSB).WithProcessor(processors.MD5)
	
	err := encode(image, &hiddenMessage)
	if err != nil {
		panic(err.Error())
	}
	
	// Then do whatever you want with the image
}

```

[1]: https://en.wikipedia.org/wiki/Steganography
