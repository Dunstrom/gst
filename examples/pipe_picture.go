package main

import "C"
import (
	"fmt"
	"github.com/Dunstrom/gst"
	"github.com/ziutek/glib"
	"os"
	"time"
)

func takePicture(srcPipeline *gst.Pipeline, filename string) {
	fmt.Fprintln(os.Stdout, "Setting up a picture pipeline")
	picturePipeline, err := gst.ParseLaunch(fmt.Sprintf("appsrc name=appsrc max-bytes=0 ! video/x-raw, format=(string)RGB, width=(int)320, height=(int)240, framerate=(fraction)30/1, multiview-mode=(string)mono, pixel-aspect-ratio=(fraction)1/1, interlace-mode=(string)progressive  ! videoconvert ! pngenc snapshot=1 ! filesink location=%s", filename))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create pipeline: ", err)
		return
	}
	fakeSink := &gst.Sink{srcPipeline.GetByName("fakesink")}
	appSrc := &gst.AppSrc{picturePipeline.GetByName("appsrc")}
	if picturePipeline.SetState(gst.STATE_PLAYING) == gst.STATE_CHANGE_FAILURE {
		fmt.Fprintln(os.Stderr, "Failed to start picture pipeline")
		return
	}

	// Start the picture pipeline
	fmt.Fprintln(os.Stdout, "Started the picture pipeline")
	picturePipeline.SetState(gst.STATE_PLAYING)

	for {
		// Pull sample
		fmt.Fprintln(os.Stdout, "Pulling a sample")
		sample := fakeSink.GetLastSample()
		if sample == nil {
			fmt.Fprintln(os.Stderr, "Failed to pull sample from fakesink")
			return
		}
		fmt.Fprintf(os.Stdout, "Sample caps: %s \n", sample.GetCaps().String())
		bufferList := sample.GetBufferList()
		if bufferList != nil {
			fmt.Fprintf(os.Stdout, "Found sample buffer list with length: %d in sample\n", bufferList.Length())
		}
		buffer := sample.GetBuffer().DeepCopy()
		fmt.Fprintf(os.Stdout, "Size of buffer in sample %v\n", buffer.GetSize())

		// Push sample
		fmt.Fprintln(os.Stdout, "Pushing the sample")
		res := appSrc.PushBuffer(buffer)
		if res == gst.GST_FLOW_EOS || res == gst.GST_FLOW_OK {
			break
		} else if res == gst.GST_FLOW_FLUSHING {
			time.Sleep(time.Millisecond * 100)
		} else if res == gst.GST_FLOW_ERROR {
			fmt.Fprintln(os.Stderr, "Error pushing the sample")
			return
		} else {
			fmt.Fprintf(os.Stderr, "Unknown return from PushBuffer %d \n", res)
		}
	}

	// Done
	fmt.Fprintf(os.Stdout, "Took picture of pipeline stored at %s \n", filename)
}

func main() {
	fmt.Fprintln(os.Stdout, "Setting up source pipeline")
	srcPipeline, err := gst.ParseLaunch("videotestsrc ! tee name=t ! queue ! autovideosink t. ! queue ! fakesink name=fakesink enable-last-sample=1") // video/x-raw, framerate=(fraction)5/1 ! fakesink name=fakesink enable-last-sample=1") // t. ! queue ! videoflip method=clockwise ! autovideosink") //")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create source pipeline: ", err)
		os.Exit(1)
	}
	if srcPipeline.SetState(gst.STATE_PLAYING) == gst.STATE_CHANGE_FAILURE {
		fmt.Fprintln(os.Stderr, "Failed to start source pipeline")
		os.Exit(1)
	}
	time.Sleep(time.Second)
	go takePicture(srcPipeline, "picture.png")
	glib.NewMainLoop(nil).Run()
}
