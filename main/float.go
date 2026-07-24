package main

import "fmt"

func anjina(a, b float64) (float64, float64) {
	area := a * b
	perimeter := 2 * (a + b)
	return area, perimeter
}
func main() {
	area, perimeter := anjina(5.5, 6.5)
	fmt.Println("Area =", area)
	fmt.Println("perimeter=", perimeter)
}
