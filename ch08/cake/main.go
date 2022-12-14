package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Shop struct {
	Verbose        bool
	Cakes          int
	BakeTime       time.Duration // time to bake one cake
	BakeStdDev     time.Duration // standard deviation of baking time
	BakeBuf        int           // buffer slots between baking and icing
	NumIcers       int           // number of cooks doing icing
	IceTime        time.Duration
	IceStdDev      time.Duration
	IceBuf         int
	InscribeTime   time.Duration
	InscribeStdDev time.Duration
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i) // nth cake to bake
		if s.Verbose {
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev) // cake need to bake and deliver
		baked <- c                     // Send baked cake to next cook
	}
	close(baked) // All cake baked, we need to close the window
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {

}

func (s *Shop) inscriber(iced <-chan cake) {

}

// work blocks the calling goroutine for a period of time
// that is normally distributed around d
// with a standard deviation of stddev.
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}

func main() {

}
