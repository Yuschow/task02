package main

import (
	"fmt"
	"sync"
	"time"
)

func odd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("odd:", i)
		}
	}
}

func even(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even:", i)
		}
	}
}

func worker(idx int, task func(), wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	fmt.Printf("task%d start\n", idx)
	task()
	elapsed := time.Since(start)
	fmt.Printf("task%d execution time:%s \n", idx, elapsed)
}

func runTask(tasks []func()) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for i, t := range tasks {
		go worker(i, t, &wg)
	}
	wg.Wait()
	fmt.Printf("All Task Done!\n")
}
