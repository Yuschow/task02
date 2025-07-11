package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	num := 10
	change(&num)
	fmt.Println("num:", num)
	change(&num)
	fmt.Println("num:", num)
	slice := []int{1, 2, 3, 4, 5}
	modifySlice(&slice)
	fmt.Println("slice:", slice)
	var wg sync.WaitGroup
	wg.Add(2)
	go odd(&wg)
	go even(&wg)
	wg.Wait()
	fmt.Print("All workers done")
	tasks := []func(){
		func() { time.Sleep(1 * time.Second) },
		func() { time.Sleep(1 * time.Second) },
		func() { time.Sleep(2 * time.Second) },
		func() { time.Sleep(2 * time.Second) },
		func() { time.Sleep(3 * time.Second) },
		func() { time.Sleep(4 * time.Second) },
	}
	runTask(tasks)
}
