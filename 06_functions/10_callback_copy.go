package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// summarize
	s := sum(ii...)
	fmt.Println(s)

	// even
	sum_even, even_list := even(sum, ii...)
	fmt.Println(even_list)
	fmt.Println(sum_even)

	// odd
	sum_odd, odd_list := odd(sum, ii...)
	fmt.Println(odd_list)
	fmt.Println(sum_odd)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) (int, []int) {
	even_list := []int{}
	for _, v := range vi {
		if v%2 == 0 {
			even_list = append(even_list, v)
		}
	}
	return f(even_list...), even_list
}

func odd(f func(xi ...int) int, vi ...int) (int, []int) {
	odd_list := []int{}
	for _, v := range vi {
		if v%2 == 1 {
			odd_list = append(odd_list, v)
		}
	}
	return f(odd_list...), odd_list
}
