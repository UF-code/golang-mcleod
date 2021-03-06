package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}
func (p person) speak() {
	fmt.Println("I am ", p.first, p.last)
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I called human")
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa2)

	sa1.speak()
	sa2.speak()
	p1.speak()

	fmt.Println("*************************************************")

	bar(sa1)
	bar(sa2)
	bar(p1)

	//
	// conversion

}
