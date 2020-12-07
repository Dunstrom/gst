package gst

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetFakesinkSamplePointer(t *testing.T) {
	pipeline, err := ParseLaunch("filesrc location=images/logo.png ! decodebin ! imagefreeze ! tee name=t ! autovideosink t. ! queue ! fakesink name=fakesink enable-last-sample=1")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create pipeline: ", err)
		os.Exit(1)
	}
	if pipeline.SetState(STATE_PLAYING) == STATE_CHANGE_FAILURE {
		fmt.Fprintln(os.Stderr, "Failed to start pipeline")
		os.Exit(1)
	}
	fakeSink := pipeline.GetByName("fakesink")
	sample := (*Sample)(fakeSink.GetPropertyPointer("last-sample"))
	assert.Greater(t, sample.GetBuffer().GetSize(), 0, "No data in sample")
}
