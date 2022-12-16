package main

import "fmt"

func main() {
	c := make(chan int, 0)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// // for need to check if channel closed
	// for {
	// 	v, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }

	// for range can break automatically once channel closed
	for i := range c {
		fmt.Println(i)
	}
}
