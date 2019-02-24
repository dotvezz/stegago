package coding

import (
	"stegago/engines"
	"stegago/processing"
	"testing"
)

var lsbEncoder = NewDecoder(engines.LSB)

func TestNew(t *testing.T) {
	_ = NewDecoder(engines.LSB)
}

func TestEncoder_WithProcessor(t *testing.T) {
	_ = lsbEncoder.WithProcessor(processing.MD5)
}
