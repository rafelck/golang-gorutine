package golang_gorutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "rafel"
		fmt.Println("Selesai mengirim data")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Rafelino Claudius Kelen"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	defer close(channel)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "rafelino"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	defer close(channel)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "rafel"
	channel <- "claudius"
	channel <- "kelen"

	fmt.Println(cap(channel)) //melihat panjang buffer
	fmt.Println(len(channel)) // melihat jumlah data di dalam buffer
}
