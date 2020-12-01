package gst

/*
#include <stdlib.h>
#include <gst/gstbuffer.h>

void _golang_gst_set_dts( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_DTS(buffer) = value;
}

void _golang_gst_set_pts( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_PTS(buffer) = value;
}

void _golang_gst_set_duration( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_DURATION(buffer) = value;
}

void _golang_gst_set_offset( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_OFFSET(buffer) = value;
}

void _golang_gst_set_offset_end( GstBuffer *buffer, guint64 value ) {
  GST_BUFFER_OFFSET_END(buffer) = value;
}

*/
import "C"

import (
	"unsafe"
)

type GstBufferStruct C.GstBuffer

type Buffer struct {
	GstBuffer *GstBufferStruct
}

func NewBuffer() *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_buffer_new())
	return buffer
}

func NewBufferAllocate(size uint) *Buffer {
	buffer := new(Buffer)
	buffer.GstBuffer = (*GstBufferStruct)(C.gst_buffer_new_allocate(nil, C.gsize(size), nil))
	return buffer
}

func (this *Buffer) g() *C.GstBuffer {
	return (*C.GstBuffer)(this.GstBuffer)
}

func (this *Buffer) GetSize() uint {
	return (uint)(C.gst_buffer_get_size((*C.GstBuffer)(this.GstBuffer)))
}

func (this *Buffer) AppendMemory(memory *Memory) {
	C.gst_buffer_append_memory((*C.GstBuffer)(this.GstBuffer), (*C.GstMemory)(memory))
}

func (this *Buffer) MemSet(offset uint, val byte, size uint) int {
	return (int)(C.gst_buffer_memset((*C.GstBuffer)(this.GstBuffer), C.gsize(offset), C.guint8(val), C.gsize(size)))
}

func (this *Buffer) FillWithGoSlice(data []byte) int {
	dataLength := uint(len(data))
	return (int)(C.gst_buffer_fill((*C.GstBuffer)(this.GstBuffer), C.gsize(0), (C.gconstpointer)(C.CBytes(data)), C.gsize(dataLength)))
}

func (this *Buffer) ExtractAll() []byte {
	var data []byte
	p := C.CBytes(data)
	defer C.free(p)
	extracted := C.gst_buffer_extract(this.g(), C.gsize(0), (C.gpointer)(p), C.gsize(this.GetSize()))
	return C.GoBytes(p, (C.int)(int(extracted)))
}

func (this *Buffer) Fill(offset uint, src unsafe.Pointer, size uint) int {
	return (int)(C.gst_buffer_fill((*C.GstBuffer)(this.GstBuffer), C.gsize(offset), C.gconstpointer(src), C.gsize(size)))
}

func (this *Buffer) SetDTS(value uint64) {
	C._golang_gst_set_dts((*C.GstBuffer)(this.GstBuffer), C.guint64(value))
}

func (this *Buffer) SetPTS(value uint64) {
	C._golang_gst_set_pts((*C.GstBuffer)(this.GstBuffer), C.guint64(value))
}

func (this *Buffer) SetDuration(value uint64) {
	C._golang_gst_set_duration((*C.GstBuffer)(this.GstBuffer), C.guint64(value))
}

func (this *Buffer) SetOffset(value uint64) {
	C._golang_gst_set_offset((*C.GstBuffer)(this.GstBuffer), C.guint64(value))
}

func (this *Buffer) SetOffsetEnd(value uint64) {
	C._golang_gst_set_offset_end((*C.GstBuffer)(this.GstBuffer), C.guint64(value))
}

func (this *Buffer) Unref() {
	C.gst_buffer_unref((*C.GstBuffer)(this.GstBuffer))
}
