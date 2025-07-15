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

	// Rectangle
	fmt.Printf("\nRectangle:\n")
	rect := Rectangle{3, 4}
	PrintShapeInfo(rect)

	// Circle
	fmt.Printf("\nCircle:\n")
	circle := Circle{3}
	PrintShapeInfo(circle)

	// Employee
	fmt.Printf("\nEmployee:\n")
	emp := Employee{
		Person:     Person{"jack", 25},
		EmployeeID: "0001",
	}
	emp.Printlnfo()
	// producer
	ch := make(chan int)
	go producer(ch)
	comsumer(ch)
	// producerBuffer
	chBuffer := make(chan int, 5)
	go producerBuffer(chBuffer)
	comsumerBuffer(chBuffer)
}
