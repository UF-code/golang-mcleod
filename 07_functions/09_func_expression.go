package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("My first func expression")
	}

	f()

	t := func(x int) {
		fmt.Println(x)
	}
	t(19)
}
