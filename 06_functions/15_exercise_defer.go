package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("First of all")
}

func foo() {
	defer func() {
		fmt.Println("defer defer foo")
	}()
	fmt.Println("defer foo")
}
