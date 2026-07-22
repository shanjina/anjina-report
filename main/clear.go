package main

import "fmt"

func resee(deepa *int) {
	*deepa = 0
}
func main() {
	anjina := 10
	resee(&anjina)
	fmt.Println("this is new one", anjina)
}
