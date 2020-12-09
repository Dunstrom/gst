package gst

import (
	"fmt"
	"unsafe"
)

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-app-1.0
#cgo LDFLAGS: -lgstapp-1.0
#include <stdlib.h>
#include <string.h>
#include <gst/gst.h>
#include <gst/app/gstappsrc.h>
*/
import "C"

type FlowReturn C.GstFlowReturn

// Read more about flow returns here https://gstreamer.freedesktop.org/documentation/gstreamer/gstpad.html?gi-language=c#GstFlowReturn
const (
	GST_FLOW_OK             = FlowReturn(C.GST_FLOW_OK)
	GST_FLOW_FLUSHING       = FlowReturn(C.GST_FLOW_FLUSHING)
	GST_FLOW_NOT_LINKED     = FlowReturn(C.GST_FLOW_NOT_LINKED)
	GST_FLOW_NOT_NEGOTIATED = FlowReturn(C.GST_FLOW_NOT_NEGOTIATED)
	GST_FLOW_ERROR          = FlowReturn(C.GST_FLOW_ERROR)
	GST_FLOW_NOT_SUPPORTED  = FlowReturn(C.GST_FLOW_NOT_SUPPORTED)
)

func (f FlowReturn) String() string {
	switch f {
	case GST_FLOW_OK:
		return "GST_FLOW_OK"
	case GST_FLOW_FLUSHING:
		return "GST_FLOW_FLUSHING"
	case GST_FLOW_NOT_LINKED:
		return "GST_FLOW_NOT_LINKED"
	case GST_FLOW_NOT_NEGOTIATED:
		return "GST_FLOW_NOT_NEGOTIATED"
	case GST_FLOW_ERROR:
		return "GST_FLOW_ERROR"
	case GST_FLOW_NOT_SUPPORTED:
		return "GST_FLOW_NOT_SUPPORTED"
	default:
		return fmt.Sprintf("flow error: %d", f)
	}
}

type AppSrc struct {
	*Element
}

func NewAppSrc(name string) *AppSrc {
	return &AppSrc{ElementFactoryMake("appsrc", name)}
}

func (a *AppSrc) g() *C.GstAppSrc {
	return (*C.GstAppSrc)(a.GetPtr())
}

func (a *AppSrc) SetCaps(caps *Caps) {
	p := unsafe.Pointer(caps) // HACK
	C.gst_app_src_set_caps(a.g(), (*C.GstCaps)(p))
}

func (a *AppSrc) EOS() error {
	ret := FlowReturn(C.gst_app_src_end_of_stream(a.g()))
	if ret != GST_FLOW_OK {
		return fmt.Errorf("appsrc eos: %v", ret)
	}

	return nil
}

func (b *AppSrc) PushBuffer(buffer *Buffer) FlowReturn {
	return (FlowReturn)(C.gst_app_src_push_buffer((*C.GstAppSrc)(b.g()), (*C.GstBuffer)(buffer.GstBuffer)))
}

func (a *AppSrc) Write(d []byte) (int, error) {
	buf := C.gst_buffer_new_allocate(nil, C.gsize(len(d)), nil)
	n := C.gst_buffer_fill(buf, C.gsize(0), (C.gconstpointer)(C.CBytes(d)), C.gsize(len(d)))

	ret := FlowReturn(C.gst_app_src_push_buffer((*C.GstAppSrc)(a.GetPtr()), buf))
	if ret != GST_FLOW_OK {
		return 0, fmt.Errorf("appsrc push buffer failed: %v", ret)
	}

	return int(n), nil
}
