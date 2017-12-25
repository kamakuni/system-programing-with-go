package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open("/")
	if err != nil {
		panic(err)
	}
	fileinfos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
	for _, fileinfo := range fileinfos {
		if fileinfo.IsDir() {
			fmt.Printf("[Dir] %s\n", fileinfo.Name())
		} else {
			fmt.Printf("[File] %s\n", fileinfo.Name())
		}
	}
}
