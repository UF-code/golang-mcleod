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
		// even if is not true if you use fallthrough it will print
		fallthrough
	case true:
		fmt.Println("This is also gonna print")
	default:
		fmt.Println("Default")
	}

	switch "Bond" {
	case "Money":
		fmt.Println("Not gonna show")
	case "Power":
		fmt.Println("Will not appear")
	case "Bond":
		fmt.Println("This will appear")
	default:
		fmt.Println("Test")
	}
}
