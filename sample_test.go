package gst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSample(t *testing.T) {
	a := assert.New(t)
	sample := NewSample(NewBuffer(), NewCapsEmpty())
	a.NotNil(sample, "Sample should not be nil")
}

func TestGetBuffer(t *testing.T) {
	a := assert.New(t)
	buffer := NewBuffer()
	data := []byte{1, 1, 1}
	sample := NewSample(buffer, NewCapsEmpty())
	buffer = sample.GetBuffer()
	dataFromBuffer := buffer.ExtractAll()
	a.Equal(data, dataFromBuffer, "Data read from sample not the same as what was written")
}
