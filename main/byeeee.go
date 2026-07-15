package main

import "fmt"

func main() {
	a := "monday",
	switch a {
	case "sunday","saturday":
		fmt.Println("Today is holiday")
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("Today is not holiday")
	}

}
