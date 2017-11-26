package main

import (
	"io"
	"os"
)

func main() {
	src, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	dst, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(dst, src)
}
