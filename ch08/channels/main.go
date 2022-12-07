package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hi")

	myCh := make(chan int, 5)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {

		fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		close(ch)
		// ch <- 5
		// ch <- 6

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
