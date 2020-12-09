package gst

/*
#include <gst/gstsample.h>
*/
import "C"

type Sample C.GstSample

func NewSample(buffer *Buffer, caps *Caps) *Sample {
	return (*Sample)(C.gst_sample_new(
		buffer.g(),
		caps.g(),
		nil, // Not implemented
		nil, // Not implemented
	))
}

func (s *Sample) g() *C.GstSample {
	return (*C.GstSample)(s)
}

func (s *Sample) GetBuffer() *Buffer {
	return &Buffer{(*GstBufferStruct)(C.gst_sample_get_buffer(s.g()))}
}

func (s *Sample) Unref() {
	C.gst_sample_unref(s.g())
}

func (s *Sample) GetCaps() *Caps {
	return (*Caps)(C.gst_sample_get_caps(s.g()))
}
