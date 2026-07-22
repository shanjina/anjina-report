package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func main() {
	x := 10
	y := 20
	fmt.Println("before", x, y)
	swap(&x, &y)
	fmt.Println("after", x, y)
}
