package gst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBufferList(t *testing.T) {
	data := [][]byte{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	bufferList := NewBufferList()
	for i, d := range data {
		buffer := NewBufferAllocate(uint(len(d)))
		buffer.FillWithGoSlice(d)
		bufferList.InsertBuffer(i, buffer)
	}
	length := bufferList.Length()
	assert.Equal(t, uint(len(data)), length, "Buffer list not of correct length")
	assert.Equal(t, uint(9), bufferList.CalculateTotalSize(), "Buffer list not of expected size")
	for i := uint(0); i < length; i += 1 {
		buffer := bufferList.GetBufferAt(i)
		assert.Equal(t, data[i], buffer.ExtractAll(), "Unexpected content in buffer")
	}
	bufferList.Remove(0, 2)
	assert.Equal(t, uint(1), bufferList.Length(), "Buffer list not of correct length")
	assert.Equal(t, data[2], bufferList.GetBufferAt(0).ExtractAll(), "Unexpected content in buffer after remove")
	bufferList.Unref()
}
