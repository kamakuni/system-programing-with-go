package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := filepath.Join(os.TempDir(), "temp.txt")
	fmt.Printf("Temp FIle Path: %s\n", path)
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)
	fragments := strings.Split(path, string(filepath.Separator))
	for _, fragment := range fragments {
		fmt.Println(fragment)
	}
}
