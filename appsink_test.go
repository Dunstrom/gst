package gst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppSink(t *testing.T) {
	a := assert.New(t)
	pipeline, err := ParseLaunch("videotestsrc pattern=snow ! video/x-raw,width=1280,height=720 ! appsink name=s")
	if err != nil {
		t.Logf("Error creating pipeline: %s", err.Error())
		t.FailNow()
	}
	_appSink := pipeline.GetByName("s")
	a.NotNil(_appSink, "Failed to get appsink")
	appSink := AppSink{_appSink}
	a.NotEqual(STATE_CHANGE_FAILURE, pipeline.SetState(STATE_PLAYING), "Failed to start pipeline")
	sample := appSink.PullSample()
	a.NotNil(sample, "Failed to get sample")
	buffer := sample.GetBuffer()
	a.NotNil(buffer, "Failed to get buffer")
	data := buffer.ExtractAll(10000)
	a.Greater(buffer.GetSize(), uint(0), "Size of buffer is not larger than 0")
	a.Equal(buffer.GetSize(), uint(len(data)), "Size of buffer and data extracted does not match")
	pipeline.SetState(STATE_NULL)
}
