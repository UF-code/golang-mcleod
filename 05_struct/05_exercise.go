package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {

	p1 := person{
		first:    "James",
		last:     "Bond",
		icecream: []string{"Almond", "Vanilla", "Strawberry"},
	}
	fmt.Println(p1)

	for i, v := range p1.icecream {
		fmt.Println(i, v)
	}
}
