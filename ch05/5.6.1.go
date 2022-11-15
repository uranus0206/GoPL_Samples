package main

import "os"

func main() {
	var rmdirs []func()

	for _, d := range tmpDirs() {
		dir := d
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	// NOTE: Incorrect way!!
	for _, d := range tmpDirs() {
		os.MkdirAll(d, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(d) // d cannot be captured
		})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}

}

func tmpDirs() []string {
	return []string{
		"AA",
	}
}
