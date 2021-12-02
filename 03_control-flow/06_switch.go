package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Hello World!")

	case (2 == 4):
		fmt.Println("This is not gonna print")
	case (2 == 2):
		fmt.Println("This is gonna print")
	case true:
		fmt.Println("This is not gonna print")
	}
}
