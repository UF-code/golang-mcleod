package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func bar() {
	fmt.Println("Not anymore")
}

func foo() {
	fmt.Println("I was the first called")
}
