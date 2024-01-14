package belajar_golang_goroutine

import (
	"fmt"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	// membuat channel dengan buffered dengan panjang 5
	channel := make(chan string, 3)

	defer close(channel)

	channel <- "eko"
	channel <- "fadly"
	channel <- "budi"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

}
