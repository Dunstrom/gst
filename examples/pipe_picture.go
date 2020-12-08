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
	picturePipeline, err := gst.ParseLaunch(fmt.Sprintf("appsrc name=appsrc num-buffers=10 ! videoconvert ! pngenc ! filesink location=%s", filename))
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
	fmt.Fprintln(os.Stdout, "Started the picture pipeline")
	time.Sleep(time.Second * 2)
	for i := 0; i < 2; i += 1 {
		fmt.Fprintln(os.Stdout, "Pulling a sample")
		sample := fakeSink.GetLastSample()
		if sample == nil {
			fmt.Fprintln(os.Stderr, "Failed to pull sample from fakesink")
			return
		}
		fmt.Fprintln(os.Stdout, "Pushing the sample")
		buffer := sample.GetBuffer().DeepCopy()
		fmt.Fprintf(os.Stdout, "Size of buffer in sample %v\n", buffer.GetSize())
		appSrc.PushBuffer(buffer)
		state, _, _ := pipeline.GetState(1000)
		if state != gst.STATE_PLAYING {
			break
		}
	}
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
