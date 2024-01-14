package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestSelectChannel(t *testing.T) {
	channel_1 := make(chan string)
	channel_2 := make(chan string)
	defer close(channel_1)
	defer close(channel_2)

	go GiveMeResponseChannel(channel_1)
	go GiveMeResponseChannel(channel_2)

	counter := 0
	for {
		select {
		case data := <-channel_1:
			fmt.Println("Print Channel 1", data)
			counter++
		case data := <-channel_2:
			fmt.Println("Print Channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func GiveMeResponseChannel(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "fadly"
}
