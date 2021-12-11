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
	p2 := person{
		first:    "Miss",
		last:     "Moneypenny",
		icecream: []string{"chocolate", "martini", "rum and coke"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(p1)

	for i, v := range p1.icecream {
		fmt.Println(i, v)
	}

	for k, v := range m {
		fmt.Println(k, v)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.icecream {
			fmt.Println(i, val)
		}
	}
}
