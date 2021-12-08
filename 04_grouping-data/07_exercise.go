package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for i, v := range nums {
		println(i, v)
	}
	fmt.Printf(`%T `, nums)

}
