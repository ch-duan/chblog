package meTest

import (
	"chblog/go/interview"
	"testing"
)

func TestSlicePointer(t *testing.T) {
	interview.SliceTestPointer()
}

func TestChannelClosed(t *testing.T) {
	interview.ChannelTestReceiveFromClosed()
}

func TestSliceLen(t *testing.T) {
	interview.SliceTestLenAndCap()
}
