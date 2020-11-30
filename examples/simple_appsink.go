package main

import (
	"fmt"
	"github.com/dunstrom/gst"
	"github.com/lijo-jose/glib"
	"os"
)

func checkElem(e *gst.Element, name string) {
	if e == nil {
		fmt.Fprintln(os.Stderr, "can't make element: ", name)
		os.Exit(1)
	}
}

func main() {
	var filename string
	if len(os.Args) > 1 && os.Args[1] != "" {
		filename = os.Args[1]
	} else {
		filename = "images/logo.png"
	}
	_, _ = fmt.Fprintf(os.Stdout, "Using file: %s as freeze frame\n", filename)
	pipeline, err := gst.ParseLaunch(fmt.Sprintf("filesrc location=%s ! decodebin ! imagefreeze ! appsink", filename))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "Failed to create pipeline because of: %s\n", err.Error())
		os.Exit(1)
	}
	_sink := pipeline.GetByName("sink")
	sink := &gst.AppSink{_sink}
	_ = pipeline.SetState(gst.STATE_PLAYING)
	glib.NewMainLoop(nil).Run()
}
