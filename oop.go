package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * 3.1415926
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.1415926
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) Printlnfo() {
	fmt.Printf("name: %s\n", e.Name)
	fmt.Printf("age: %d\n", e.Age)
	fmt.Printf("EmployeeID: %s\n", e.EmployeeID)
}
