package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("pid %d\n", os.Getegid())
	fmt.Printf("parent pid %d\n", os.Getppid())
}
