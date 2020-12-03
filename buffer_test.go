package gst

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBufferExtract(t *testing.T) {
	a := assert.New(t)
	data := make([]byte, 10000)
	rand.Read(data)
	buffer := NewBufferAllocate(uint(len(data)))
	a.Equal(len(data), buffer.FillWithGoSlice(data), "All of the data was not written to the buffer")
	a.Equal(data, buffer.ExtractAll(10000), "Filled and extracted data not the same")
	buffer.Unref()
}
