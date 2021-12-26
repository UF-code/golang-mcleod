package main

import "fmt"

func main() {
	r := factorial_recursion(4)
	fmt.Println(r)

	l := factorial_loop(4)
	fmt.Println(l)

}

func factorial_recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial_recursion(n-1)
}

func factorial_loop(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}
