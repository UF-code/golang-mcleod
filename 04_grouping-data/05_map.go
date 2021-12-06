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
	// map if there is item
	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}

	//
	// map add
	m["Todd"] = 43

	for k, v := range m {
		println(k, v)
	}

	//
	// map delete

	if v, ok := m["Ian Fleming"]; ok {
		fmt.Println(v)
		delete(m, "James")
	}
}
