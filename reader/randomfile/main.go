package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	dst, err := os.Create("random.txt")
	if err != nil {
		panic(err)
	}
	io.CopyN(dst, rand.Reader, 1024)
}
