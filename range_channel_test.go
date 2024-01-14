package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke" + strconv.Itoa(i)

		}
		close(channel)
	}()

	// iterasi didalam range channel
	for value := range channel {
		fmt.Println("Value:", value)
	}
	fmt.Println("Done")
}
