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
	defer src.Close()
	dst, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	io.Copy(dst, src)
}
