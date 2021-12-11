package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)

	type foo int
	var y foo

	y = 42
	fmt.Printf("%T", y)
}
