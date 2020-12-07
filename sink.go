package gst

import (
	"github.com/ziutek/glib"
	"unsafe"
)

type Sink struct {
	*Element
}

func (s *Sink) GetLastSample() *Sample {
	return (*Sample)(unsafe.Pointer(s.GetProperty("last-sample").(glib.Pointer)))
}
