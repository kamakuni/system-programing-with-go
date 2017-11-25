package main

import (
	"bufio"
	"os"
)

func main() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.Write([]byte("bufio.Writer "))
	buffer.Flush()
	buffer.Write([]byte("example \n"))
	buffer.Flush()
}
