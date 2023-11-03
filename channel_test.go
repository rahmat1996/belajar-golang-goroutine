package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // make channel
	defer close(channel)         // close channel

	// anonymous function
	go func() {
		time.Sleep(2 * time.Second) // wait 2 seconds
		channel <- "Rahmat Belajar Golang"
		fmt.Println("Send data to channel done")
	}()

	data := <-channel // send data channel to data
	fmt.Println(data) // print data

	time.Sleep(5 * time.Second) // wait 5 seconds
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Rahmat Belajar Golang"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string) // make channel
	defer close(channel)         // close channel

	go GiveMeResponse(channel)

	data := <-channel // send data channel to data
	fmt.Println(data) // print data

	time.Sleep(5 * time.Second) // wait 5 seconds
}

// function just send data to channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Rahmat Belajar Golang"
}

// function just get data from channel
func OnlyOut(channel <-chan string) {
	data := <-channel // send data channel to data
	fmt.Println(data) // print data
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string) // make channel
	defer close(channel)         // close channel

	go OnlyIn(channel)  // function send data to channel
	go OnlyOut(channel) // funtion get data from channel

	time.Sleep(5 * time.Second) // wait 5 seconds
}

// buffered channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Rahmat"
		channel <- "Belajar"
		channel <- "Golang"
		// channel <- "Programming" // if more 3 goroutine will deadlock!
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Done")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- "Loop " + strconv.Itoa(i)
		}
		close(channel) // channel must close or will deadlock (btw if use defer it will deadlock)
	}()

	for data := range channel {
		fmt.Println("Get data", data)
	}

	fmt.Println("Done")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		default:
			fmt.Println("Please waiting.") // this code will execute if channel 1 and channel 2 not send data or late send data
		}

		if counter == 2 {
			break
		}
	}
}
