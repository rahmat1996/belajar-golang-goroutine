package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // call in end execution will decrease group.Add

	group.Add(1) // there 1 progress must be waiting

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}
	// running goroutine 100
	for i := 1; i <= 100; i++ {
		go RunAsynchronous(group)
	}
	group.Wait() // waiting all goroutine already done
	fmt.Println("Done")
}
