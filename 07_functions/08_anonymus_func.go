package main

import "fmt"

func foo() {
	fmt.Println("Foo run")
}

func main() {
	foo()

	func() {
		fmt.Println("Anonymus func run")
	}()

	func(x int) {
		fmt.Printf("Number is %v \n", x)
	}(42)

}
