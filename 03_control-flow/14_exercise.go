package main

import "fmt"

func main() {

	switch "Basketball" {
	case "football":
		fmt.Println("This sport for suckers")
	case "Basketball":
		fmt.Println("This sport for man!")
	default:
		fmt.Println("your sport is not qualified for me")
	}

}
