package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {

	group := sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	// Total CPU
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	// Total Thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	// Total Goroutine
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {

	group := sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	// Total CPU
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	// change thread number
	runtime.GOMAXPROCS(20)
	// Total Thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	// Total Goroutine
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	group.Wait()
}
