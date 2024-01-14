package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// membuat channel
	channel := make(chan string)

	// mmembuat goroutine
	go func() {
		time.Sleep(10 * time.Second)
		channel <- "fadly" // mengirim data ke channel
	}()
	// menunggu sampai ada data di channel
	data := <-channel
	fmt.Println("Data diterima : ", data)

	// menutup channel
	close(channel)

}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	// membuat goroutine
	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println("Data : ", data)
	close(channel)
}

// channel menggunakan parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "fadly"
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	// membuat goroutine
	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(4 * time.Second)
	close(channel)

}

// membuat channel sebagai penerima
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Channel Test In"
}

// membuat channel sebagai pengirim
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}
