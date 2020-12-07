package main

import "C"
import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/Dunstrom/gst"
	"github.com/ziutek/glib"
)

func main() {
	pipeline, err := gst.ParseLaunch("filesrc location=images/logo.png ! decodebin ! imagefreeze ! tee name=t ! autovideosink t. ! queue ! fakesink name=fakesink enable-last-sample=1")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create pipeline: ", err)
		os.Exit(1)
	}
	if pipeline.SetState(gst.STATE_PLAYING) == gst.STATE_CHANGE_FAILURE {
		fmt.Fprintln(os.Stderr, "Failed to start pipeline")
		os.Exit(1)
	}

	go func() {
		i := 0
		for {
			time.Sleep(time.Second * 2)
			filename := fmt.Sprintf("pipe_picture_%d.jpg", i)
			picturePipeline, err := gst.ParseLaunch(fmt.Sprintf("appsrc name=appsrc ! jpegenc ! filesink location=%s", filename))
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to create pipeline: ", err)
				os.Exit(1)
			}
			fakeSink := pipeline.GetByName("fakesink")
			//appSrc := &gst.AppSrc{picturePipeline.GetByName("appsrc")}
			if picturePipeline.SetState(gst.STATE_PLAYING) == gst.STATE_CHANGE_FAILURE {
				fmt.Fprintln(os.Stderr, "Failed to start picture pipeline")
				os.Exit(1)
			}
			sample := (*gst.Sample)(fakeSink.GetPropertyPointer("last-sample"))
			if sample == nil {
				fmt.Fprintln(os.Stderr, "Failed to pull sample from fakesink")
			}
			fmt.Fprintln(os.Stdout, "Sample type: ", reflect.TypeOf(sample))
			fmt.Fprintln(os.Stdout, "sample: ", sample)
			//appSrc.PushBuffer(sample.GetBuffer())
			i += 1
			time.Sleep(time.Second * 8)
		}
	}()
	glib.NewMainLoop(nil).Run()
}
