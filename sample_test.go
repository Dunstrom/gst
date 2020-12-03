package gst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSample(t *testing.T) {
	a := assert.New(t)
	sample := NewSample(NewBuffer(), NewCapsEmpty())
	a.NotNil(sample, "Sample should not be nil")
	sample.Unref()
}

func TestGetBuffer(t *testing.T) {
	a := assert.New(t)
	data := []byte{1, 1, 1}
	buffer := NewBufferAllocate(uint(len(data)))
	buffer.FillWithGoSlice(data)
	sample := NewSample(buffer, NewCapsEmpty())
	a.Equal(data, sample.GetBuffer().ExtractAll(), "Data read from sample not the same as what was written")
	sample.Unref()
	buffer.Unref()
}
