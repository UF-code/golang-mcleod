package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	//
	//
	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}
}
