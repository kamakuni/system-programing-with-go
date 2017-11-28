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
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
