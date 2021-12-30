package main

import "fmt"

func main() {

	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()

	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}
	fmt.Println("about to exit")

}
