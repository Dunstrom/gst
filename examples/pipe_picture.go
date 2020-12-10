package main

import "C"
import (
	"fmt"
	"github.com/Dunstrom/gst"
	"os"
	"time"
)

func takePicture(pipeline *gst.Pipeline, filename string) {
	fmt.Fprintln(os.Stdout, "Setting up a picture pipeline")
	picturePipeline, err := gst.ParseLaunch(fmt.Sprintf("appsrc name=appsrc max-bytes=0 ! videoconvert !video/x-raw, format=(string)RGB, width=(int)320, height=(int)240, framerate=(fraction)30/1, multiview-mode=(string)mono, pixel-aspect-ratio=(fraction)1/1, interlace-mode=(string)progressive ! pngenc snapshot=1 ! filesink location=%s", filename))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create pipeline: ", err)
		os.Exit(1)
	}
	fakeSink := &gst.Sink{pipeline.GetByName("fakesink")}
	appSrc := &gst.AppSrc{picturePipeline.GetByName("appsrc")}
	if picturePipeline.SetState(gst.STATE_PLAYING) == gst.STATE_CHANGE_FAILURE {
		fmt.Fprintln(os.Stderr, "Failed to start picture pipeline")
		os.Exit(1)
	}

	// Pull sample
	fmt.Fprintln(os.Stdout, "Pulling a sample")
	sample := fakeSink.GetLastSample()
	if sample == nil {
		fmt.Fprintln(os.Stderr, "Failed to pull sample from fakesink")
		return
	}
	fmt.Fprintf(os.Stdout, "Sample caps: %s\n", sample.GetCaps().String())
	bufferList := sample.GetBufferList()
	if bufferList != nil {
		fmt.Fprintf(os.Stdout, "Found sample buffer list with length: %d in sample\n", bufferList.Length())
	}
	buffer := sample.GetBuffer().DeepCopy()
	fmt.Fprintf(os.Stdout, "Size of buffer in sample %v\n", buffer.GetSize())

	// Stop src pipeline
	pipeline.SetState(gst.STATE_NULL)
	time.Sleep(time.Second * 2)
	fmt.Fprintln(os.Stdout, "Stopped pipeline")

	// Start picture pipeline
	fmt.Fprintln(os.Stdout, "Started the picture pipeline")
	picturePipeline.SetState(gst.STATE_PLAYING)
	time.Sleep(time.Second * 2)

	// Push sample
	fmt.Fprintln(os.Stdout, "Pushing the sample")
	res := appSrc.PushBuffer(buffer)
	if res != gst.GST_FLOW_OK {
		fmt.Fprintf(os.Stdout, "Failed to push buffer with error code %d\n", res)
	}
	fmt.Fprintf(os.Stdout, "Pushed sample with res %d\n", res)

	// Done
	fmt.Fprintf(os.Stdout, "Took picture of pipeline stored at %s\n", filename)
}

func main() {
	pipeline, err := gst.ParseLaunch("videotestsrc ! fakesink name=fakesink enable-last-sample=1") // video/x-raw, framerate=(fraction)5/1 ! fakesink name=fakesink enable-last-sample=1") // t. ! queue ! videoflip method=clockwise ! autovideosink") //")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create pipeline: ", err)
		os.Exit(1)
	}
	if pipeline.SetState(gst.STATE_PLAYING) == gst.STATE_CHANGE_FAILURE {
		fmt.Fprintln(os.Stderr, "Failed to start pipeline")
		os.Exit(1)
	}
	takePicture(pipeline, "picture.png")
}
