package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 1; i <= 10; i++ {
		go WaitCondition(i)
	}

	// will wait one by one
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// // will send all after wait
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
}
