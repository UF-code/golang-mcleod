package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	sum_even, even_list := even(sum, x...)
	fmt.Println(sum_even, even_list)

	sum_odd, odd_list := odd(sum, x...)
	fmt.Println(sum_odd, odd_list)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, y ...int) (int, []int) {
	even_list := []int{}

	for _, v := range y {
		if v%2 == 0 {
			even_list = append(even_list, v)
		}
	}

	return f(even_list...), even_list
}

func odd(f func(x ...int) int, y ...int) (int, []int) {
	odd_list := []int{}

	for _, v := range y {
		if v%2 == 1 {
			odd_list = append(odd_list, v)
		}
	}

	return f(odd_list...), odd_list
}
