package main

import "fmt"

func main() {

	// anonymus func
	func() {
		fmt.Println("Hey")
	}()

	// anoymus func with argument
	func(x int) {
		fmt.Printf("Value: %v Type: %T \n", x, x)
	}(52)
}
