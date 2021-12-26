package main

import "fmt"

func main() {
	foo()

	// anonymus function without parameter
	func() {
		fmt.Println("Anonymus func run")
	}()

	// anonymous function with arguments
	func(x int) {
		fmt.Println("Value of X is ", x)
	}(42)

}

func foo() {
	fmt.Println("foo ran")
}
