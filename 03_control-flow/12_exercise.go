package main

import "fmt"

func main() {
	condition := true

	if !condition {
		fmt.Println("This if not gonna show")

	} else if condition {
		fmt.Println("This else if gonna show")
	} else {
		fmt.Println("This else not gonna show")
	}
}
