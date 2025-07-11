package main

import "fmt"

func main() {
	num := 10
	change(&num)
	fmt.Println("num:", num)
	change(&num)
	fmt.Println("num:", num)
	slice := []int{1, 2, 3, 4, 5}
	modifySlice(&slice)
	fmt.Println("slice:", slice)
}
