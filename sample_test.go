package gst

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestNewSample(t *testing.T) {
	a := assert.New(t)
	sample := NewSample(NewBuffer(), NewCapsEmpty())
	a.NotNil(sample, "Sample should not be nil")
	sample.Unref()
}

func TestGetBuffer(t *testing.T) {
	a := assert.New(t)
	data := []byte{1, 1, 1}
	buffer := NewBufferAllocate(uint(len(data)))
	buffer.FillWithGoSlice(data)
	sample := NewSample(buffer, NewCapsEmpty())
	a.Equal(data, sample.GetBuffer().ExtractAll(), "Data read from sample not the same as what was written")
	sample.Unref()
	buffer.Unref()
}

func TestGetCaps(t *testing.T) {
	capsString := "video/x-raw, format=(string)AYUV64, width=(int)320, height=(int)240, framerate=(fraction)30/1, multiview-mode=(string)mono, pixel-aspect-ratio=(fraction)1/1, interlace-mode=(string)progressive"
	pipeline, err := ParseLaunch(fmt.Sprintf("videotestsrc ! %s ! fakesink name=fakesink enable-last-sample=1", capsString))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create pipeline for get caps test")
		t.FailNow()
	}
	if pipeline.SetState(STATE_PLAYING) == STATE_CHANGE_FAILURE {
		fmt.Fprintln(os.Stderr, "Failed to start pipeline for get caps test")
		t.FailNow()
	}
	fakeSink := &Sink{pipeline.GetByName("fakesink")}
	time.Sleep(time.Second * 1)
	sample := fakeSink.GetLastSample()
	caps := sample.GetCaps()
	assert.Equal(t, capsString, caps.String())
}
