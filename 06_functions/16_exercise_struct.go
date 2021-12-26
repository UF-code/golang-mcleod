package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.age)
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}
	p1.speak()
}
