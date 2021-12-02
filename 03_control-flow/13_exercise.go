package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Not gonna show")
	case true:
		fmt.Println("Gonna show")
	}
}
