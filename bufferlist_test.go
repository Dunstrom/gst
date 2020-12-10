package gst

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBufferList(t *testing.T) {
	const listLength int = 10
	const bufferLength int = 10000
	var data [listLength][bufferLength]byte
	for i := 0; i < listLength; i += 1 {
		d := make([]byte, bufferLength)
		rand.Read(d)
		copy(data[i][:], d)
	}
	bufferList := NewBufferList()
	for i, d := range data {
		buffer := NewBufferAllocate(uint(len(d)))
		buffer.FillWithGoSlice(d[:])
		bufferList.InsertBuffer(i, buffer)
	}
	length := bufferList.Length()
	assert.Equal(t, uint(len(data)), length, "Buffer list not of correct length")
	assert.Equal(t, uint(listLength*bufferLength), bufferList.CalculateTotalSize(), "Buffer list not of expected size")
	for i := uint(0); i < length; i += 1 {
		buffer := bufferList.GetBufferAt(i)
		assert.Equal(t, data[i][:], buffer.ExtractAll(), "Unexpected content in buffer")
	}
	bufferList.Remove(0, 2)
	assert.Equal(t, uint(listLength-2), bufferList.Length(), "Buffer list not of correct length")
	assert.Equal(t, data[2][:], bufferList.GetBufferAt(0).ExtractAll(), "Unexpected content in buffer after remove")
	bufferList.Unref()
}
