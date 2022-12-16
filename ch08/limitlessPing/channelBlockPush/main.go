package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)
	start := time.Now()
	defer fmt.Println(time.Since(start))

	go func() {
		fmt.Println("calculate goroutine starts calculating")
		time.Sleep(time.Second)
		fmt.Println("calculate goroutine ends calculating")

		ch <- "FINISH"
		fmt.Println("calculate goroutine finished")
	}()

	// // time.Sleep(2 * time.Second) // making main routine slower than above
	// fmt.Println(<-ch) // blocked by goroutine above!
	// time.Sleep(time.Second)
	// fmt.Println("main goroutine finished ", time.Since(start))

	// Select can prevent wait other goroutine
	select {
	case <-ch:
		fmt.Println("main goroutine finished")
		return
	default:
		fmt.Println("WAITing..")
		time.Sleep(500 * time.Millisecond)
	}
}
