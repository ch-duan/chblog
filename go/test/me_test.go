package meTest

import (
	"chblog/go/interview"
	"fmt"
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

func TestTt(t *testing.T) {
	a := 1
	fmt.Printf("0x%X\n", &a)
	fmt.Printf("%p", &a)
}
