package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with os stdout at %v", time.Now())
}
