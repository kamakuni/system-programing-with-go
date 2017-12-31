package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Match("text-*.txt", "text-100.txt"))
	files, err := filepath.Glob("./*.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}
