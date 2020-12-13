package gst

/*
#include <gst/gstsample.h>
*/
import "C"

import (
	"github.com/ziutek/glib"
	"reflect"
	"unsafe"
)

type Sink struct {
	*Element
}

func (s *Sink) GetLastSample() *Sample {
	sample := s.GetProperty("last-sample")
	if reflect.TypeOf(sample) == nil {
		return nil
	}
	return (*Sample)(unsafe.Pointer(sample.(glib.Pointer)))
}
