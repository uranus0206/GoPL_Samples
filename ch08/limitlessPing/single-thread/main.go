package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func say2(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func say3(s string, c chan interface{}) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

	c <- "FINISH"
}

func main() {
	fmt.Println("Num of CPU: ", runtime.NumCPU())

	since := time.Now()

	// wg := new(sync.WaitGroup)
	// wg.Add(2)
	// go say2("world", wg)
	// go say2("hello", wg)
	// wg.Wait()

	// go say("world")
	// go say("hello")
	// time.Sleep(10 * time.Second)

	ch := make(chan interface{}, 0)
	go say3("world", ch)
	go say3("hello", ch)

	<-ch
	<-ch

	defer fmt.Println(time.Since(since))
}
