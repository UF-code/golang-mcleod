package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond":       {"Shaken not stirred", "Martinis", "Women"},
		"moneypenny": {"James Bond", "Literature", "Computer Science"},
		"kofuz":      {"Being Evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(m)
	fmt.Println(len(m))

	m["igor"] = []string{"Stand still", "Hard as a rock", "Missing in action"}

	for k, v := range m {
		fmt.Printf(`This is the record for %v
`, k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
