package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	// int sorting
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	// string sorting
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}
