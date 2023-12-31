package belajar_golang_goroutine

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	for i := 1; i <= 1000; i++ {

		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				// x = x + 1
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()

	}

	group.Wait()
	fmt.Println("Counter =", x)
}
