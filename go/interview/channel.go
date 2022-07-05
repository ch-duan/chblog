package interview

import "fmt"

func ChannelTestReceiveFromClosed() {
	ch := make(chan string)
	ch <- "123"
	// close(ch)
	x, ok := <-ch
	if ok {
		fmt.Println("received: ", x)
	}

	x, ok = <-ch
	if !ok {
		fmt.Println("channel closed, data invalid.", x)
	}

	x, ok = <-ch
	if !ok {
		fmt.Println("channel closed, data invalid.", x)
	}
	testCh := make(chan string)
	close(testCh)
	x, ok = <-testCh
	if !ok {
		fmt.Println(" testCh channel closed, data invalid.", x)
	}
}
