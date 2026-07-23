package main

import "fmt"

func changevaluebyp(n *anjina, prabisha string) {
	(*n).name = prabisha
}
func changevaluebyp(n anjina, newvalue string) {
	(n).name = newvalue
}

type anjina struct {
	name string
	age  int
}

func main() {
	deepa := anjina{
		name: "anjina",
		age:  17,
	}
	fmt.Println(deepa.age)
	changevalue(&deepa, "ayush")
	fmt.Println(deepa.name)
	changevaluebyv(deepa, "anjina")
	fmt.Println(deepa.name)

}
