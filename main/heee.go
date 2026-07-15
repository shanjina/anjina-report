package main

import "fmt"

func main() {
	a := "Sunda"
	switch a {
	case "Sunday":
		fmt.Println("Today is Sunday")
	case "Monday":
		fmt.Println("Today is Monday")
	default:
		fmt.Println("Sorry,I didnt understand")
	}
}
