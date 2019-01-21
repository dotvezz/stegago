package encoders

import (
	"stegago/engines"
	"stegago/processors"
	"testing"
)

var lsbEncoder = New(engines.LSB)

func TestNew(t *testing.T) {
	_ = New(engines.LSB)
}

func TestEncoder_WithProcessor(t *testing.T) {
	_ = lsbEncoder.WithProcessor(processors.MD5)
}