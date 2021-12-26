package main

import "fmt"

// struct

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secret agent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// interface
// if the type has speak functionality
// it gives also human type
type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I called Human", h, " - inside bar")
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa1.speak()
	bar(sa1)

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	sa2.speak()
	bar(sa2)

	p1 := person{
		first: "Dr.",
		last:  "Mundo",
	}
	p1.speak()
	bar(p1)
}
