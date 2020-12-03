package gst

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-app-1.0
#cgo LDFLAGS: -lgstapp-1.0
#include <stdlib.h>
#include <string.h>
#include <gst/gst.h>
#include <gst/app/gstappsink.h>
*/
import "C"

type AppSink struct {
	*Element
}

func NewAppSink(name string) *AppSink {
	return &AppSink{ElementFactoryMake("appsink", name)}
}

func (a *AppSink) g() *C.GstAppSink {
	return (*C.GstAppSink)(a.GetPtr())
}

func (a *AppSink) PullSample() *Sample {
	return (*Sample)(C.gst_app_sink_pull_sample(a.g()))
}

func (a *AppSink) Read(max uint) []byte {
	sample := a.PullSample() // Blocks until something to read or EOF
	buffer := sample.GetBuffer()
	return buffer.ExtractAll(max)
}
