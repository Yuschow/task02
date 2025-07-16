package main

import (
	"sync"
	"sync/atomic"
)

func add(x *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	defer mu.Unlock()
	mu.Lock()
	for range 1000 {
		*x++
	}
}

func addAtomic(y *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 1000 {
		atomic.AddInt64(y, 1)
	}
}
