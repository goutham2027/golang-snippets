package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	n := 0
	fmt.Println(increment(&n, 10))

	var n32 int32
	fmt.Println(incrementAtomic(&n32, 10))
}

func increment(n *int, times int) int {
	var wg sync.WaitGroup

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func(n *int) {
			defer wg.Done()
			*n++
		}(n)
	}

	wg.Wait()
	return *n
}

func incrementAtomic(n *int32, times int) int32 {
	var wg sync.WaitGroup

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func(n *int32) {
			defer wg.Done()
			atomic.AddInt32(n, 1)
		}(n)
	}

	wg.Wait()
	return *n
}

func incrementMutex(n *int, times int) int {
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func(n *int) {
			defer wg.Done()
			mu.Lock()
			*n++
			mu.Unlock()
		}(n)
	}
	return *n
}
