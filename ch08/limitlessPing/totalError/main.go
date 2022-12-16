package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeNumber struct {
	v   int
	mux sync.RWMutex
}

func main() {
	since := time.Now()

	// total := 0

	// // Race condition!
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		total += 1
	// 	}()
	// }

	// time.Sleep(time.Second)
	// fmt.Println(total)

	// // Mutex lock, unlock
	// total := SafeNumber{v: 0}
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		total.mux.Lock()
	// 		total.v++
	// 		total.mux.Unlock()
	// 	}()
	// }
	// time.Sleep(time.Second)
	// total.mux.Lock()
	// fmt.Println(total.v)
	// total.mux.Unlock()

	// Channel
	total := 0
	ch := make(chan int, 1) // 1 buffer, no buffer will cause deadlock
	ch <- total             // push value to channel

	for i := 0; i < 1000; i++ {
		go func() {
			ch <- <-ch + 1 // get value from channel and push by add 1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(<-ch)

	defer fmt.Println(time.Since(since))
}
