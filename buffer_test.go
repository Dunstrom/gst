package gst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBufferExtract(t *testing.T) {
	a := assert.New(t)
	data := []byte{1, 2, 3}
	buffer := NewBufferAllocate(uint(len(data)))
	a.Equal(len(data), buffer.FillWithGoSlice(data), "All of the data was not written to the buffer")
	a.Equal(data, buffer.ExtractAll(), "Filled and extracted data not the same")
}
