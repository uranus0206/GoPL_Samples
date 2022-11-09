package main

import (
	"fmt"
	"sort"
)

var m = make(map[string]int)

func main() {

	// ages := make(map[string]int)
	// ages["alice"] = 31
	// ages["charlie"] = 34

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	// delete(ages, "alice")
	delete(ages, "bob")

	ages["bob"] += 1
	// _ = &ages["bob"] // Cannot take address of map.

	names := make([]string, 0, len(ages))
	for name, age := range ages {
		names = append(names, name)
		fmt.Printf("%s\t%d\n", name, age)
	}

	sort.Strings(names)
	fmt.Printf("Sorted names: %v\n", names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	// Assign value to nil map will cause panic
	// map need to instance
	// var nilAges map[string]int
	// nilAges["carol"] = 21

	isEqual := equal(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Println(isEqual)

	aList := []string{
		"A",
		"B",
		"C",
	}
	a := k(aList)
	fmt.Printf("%v\n", a)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
