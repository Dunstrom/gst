package gst

import (
	"fmt"
)

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

func (a *AppSink) Read() ([]byte, error) {
	var d []byte
	fmt.Println("Reading from AppSink")
	gstSample := C.gst_app_sink_pull_sample(a.g()) // Blocks until something to read or EOF
	fmt.Printf("Read sample %v from AppSink\n", gstSample)
	gstBuffer := C.gst_sample_get_buffer(gstSample)
	fmt.Printf("Got buffer %v from sample\n", gstBuffer)
	C.gst_buffer_extract(gstBuffer, C.gsize(0), (C.gpointer)(C.CBytes(d)), C.gst_buffer_get_size(gstBuffer))
	return d, nil
}
