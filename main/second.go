package main

import "fmt"

func main() {
	a := 10
	b := 3.00
	fmt.Println("Adiition:", float64(a)+b)
	fmt.Println("Subtraction:", float64(a)-b)
	fmt.Println("Multiplication:", float64(a)*b)
	fmt.Println("division:", float64(a)/b)
	fmt.Println("Modulus:", a%int(b))

}
