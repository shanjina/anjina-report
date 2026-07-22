package main

import "fmt"

func namaskar(a *string) {
	*a = "namaskar" + *a
}
func main() {
	b := "deepa"
	namaskar(&b)
	fmt.Println(b)
}
