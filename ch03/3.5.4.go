package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(basename("a/b/c.go"))
	// fmt.Println(basename2("a/b/c.go"))
	// fmt.Println(intsToString([]int{1, 2, 3}))
	fmt.Println(comma("123456789"))
	fmt.Println(comma2("123456789"))
}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// intsToString is linke fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func comma2(s string) string {
	var buf bytes.Buffer
	n := len(s)

	if n > 3 {
		for i := 0; i < n; i++ {
			if i%3 == 0 && i != 0 {
				buf.WriteString(",")
			}
			fmt.Fprintf(&buf, "%c", s[i])
		}
		return buf.String()
	} else {
		return s
	}
}
