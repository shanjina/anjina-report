package main

import "fmt"

func togglebool(a *bool) {
	*a = !*a
}
func main() {
	c := true
	fmt.Println("before:", c)
	togglebool(&c)
	fmt.Println("after:", c)
}
