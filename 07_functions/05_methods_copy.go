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

// func (r receiver) identifier(parameters) (return(s)) {code}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "and I have license to kill:", s.ltk)
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

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	sa2.speak()
}
