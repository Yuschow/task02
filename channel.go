package main

import "fmt"

func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close((ch))
}

func comsumer(ch <-chan int) {
	for v := range ch {
		fmt.Printf("message: %d\n", v)
	}
}

func producerBuffer(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		if len(ch) == cap(ch) {
			fmt.Println("channel is full")
		}
		ch <- i
	}
	close((ch))
}

func comsumerBuffer(ch <-chan int) {
	for v := range ch {
		fmt.Printf("message: %d\n", v)
	}
}
