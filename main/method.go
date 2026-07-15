package main

import "fmt"

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}
func main() {
	rect := Rectangle{
		length:  12,
		breadth: 6,
	}
	fmt.Println("Perimeter of Rectangle:", rect.Perimeter())

}
