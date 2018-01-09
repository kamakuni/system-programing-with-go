package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("user id: %d\n", os.Getuid())
	fmt.Printf("group id: %d\n", os.Getegid())
	groups, _ := os.Getgroups()
	fmt.Printf("groups %v\n", groups)
}
