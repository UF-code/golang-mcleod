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

//
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "and I have ltk:", s.ltk, "---secretAgent")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "---person")
}

func bar(h human) {
	switch h.(type) {
	case person:
		// asserting
		fmt.Println("I was passed into bar ---person", h.(person).first)

	case secretAgent:
		// asserting
		fmt.Println("I was passed into bar ---secretAgent", h.(secretAgent).first)
	}
	fmt.Println("I called human", h)
}

// interfaces
type human interface {
	speak()
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person{
			"Miss",
			"Moneypenny",
		},
		true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Mundo",
	}

	// secret agent 1
	fmt.Println(sa1)
	sa1.speak()
	bar(sa1)
	// secret agent 2
	fmt.Println(sa2)
	sa2.speak()
	bar(sa2)
	// person 1
	fmt.Println(p1)
	p1.speak()
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Printf("Value: %v Type: %T \n", x, x)

	z := int(x)
	fmt.Printf("Value: %v Type: %T \n", z, z)
}
