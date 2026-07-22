package main

import "fmt"

func main() {
	s := "Hello, Go!"
	p := &s

	fmt.Println("Address:", p)
	fmt.Println("Value:", *p)
}