package main

import "fmt"

func main() {
	age := 25
	var ptr *int = &age
	fmt.Println(*ptr)
	fmt.Println(age)
	fmt.Println(&age)
	fmt.Println(ptr)
}
