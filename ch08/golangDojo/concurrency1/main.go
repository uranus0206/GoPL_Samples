package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	smokeSignal := make(chan bool)
	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Andy"}

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, smokeSignal)
	}

	select {
	case <-smokeSignal:
		break
	default:
		fmt.Println("Hi")
	}
}

func attack(target string, attacked chan bool) {
	fmt.Println("Throwing ninja stars at", target)
	// time.Sleep(time.Second)
	attacked <- true
}
