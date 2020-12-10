package gst

/*
#include <stdlib.h>
#include <gst/gstbufferlist.h>
*/
import "C"

type BufferList C.GstBufferList

func (bl *BufferList) g() *C.GstBufferList {
	return (*C.GstBufferList)(bl)
}

func NewBufferList() *BufferList {
	return (*BufferList)(C.gst_buffer_list_new())
}

func (bl *BufferList) CalculateTotalSize() uint {
	return uint(C.gst_buffer_list_calculate_size(bl.g()))
}

func (bl *BufferList) GetBufferAt(idx uint) *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_buffer_list_get(bl.g(), C.guint(idx)))
	return buffer
}

func (bl *BufferList) InsertBuffer(idx int, buffer *Buffer) {
	C.gst_buffer_list_insert(bl.g(), C.gint(idx), buffer.g())
}

func (bl *BufferList) Length() uint {
	return uint(C.gst_buffer_list_length(bl.g()))
}

func (bl *BufferList) Remove(idx uint, length uint) {
	C.gst_buffer_list_remove(bl.g(), C.guint(idx), C.guint(length))
}

func (bl *BufferList) Unref() {
	C.gst_buffer_list_unref(bl.g())
}
