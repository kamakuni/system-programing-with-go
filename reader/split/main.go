package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `line 1
line 2
line 3`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err != nil {
			break
		}
	}
}
