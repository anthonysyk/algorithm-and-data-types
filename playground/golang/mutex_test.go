package golang

import (
	"fmt"
	"sync"
	"testing"
)

type MyMutex struct {
	counter int
	mu      sync.Mutex
}

func (mm *MyMutex) inc(value int) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	mm.counter += value
}

func TestMutex(t *testing.T) {
	var test MyMutex
	var wg sync.WaitGroup
	doIncrement := func(value int) {
		test.inc(value)
		wg.Done()
	}

	wg.Add(3)
	go doIncrement(1)
	go doIncrement(1)
	go doIncrement(1)
	wg.Wait()
	fmt.Println("end of test", test.counter)
}
