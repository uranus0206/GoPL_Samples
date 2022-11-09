package main

import (
	"fmt"
	"math"
	"time"
)

const (
	e       = 2.31818
	pi      = 3.124159
	IPv4Lan = 4
)

const noDelay time.Duration = 0
const timeout = 5 * time.Minute

// b directly copy value of a
// d directluy copy vcalue of c
const (
	a = 1
	b
	c = 2
	d
)

// use itoa with const group to generate const sequence
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

const Pi64 float64 = math.Pi

var x float32 = float32(Pi64)
var y float64 = Pi64
var z complex128 = complex128(Pi64)

func main() {
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	fmt.Println(a, b, c, d)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)

	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))

	fmt.Println(z)

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)
	fmt.Println(5 / 9 * (f - 32))
	fmt.Println(5.0 / 9.0 * (f - 32))
}

// netflag
func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }
