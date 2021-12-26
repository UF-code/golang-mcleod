package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	f := foo(bar, ii...)
	fmt.Println(f)

	b := bar(ii...)
	fmt.Println(b)
}

func foo(f func(xi ...int) int, vi ...int) int {
	li := []int{}
	for _, v := range vi {
		li = append(li, v)
	}
	return f(li...)
}

func bar(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
