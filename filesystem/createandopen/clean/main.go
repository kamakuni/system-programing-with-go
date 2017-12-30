package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func clean2(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = my.HomeDir + path[1:]
	}
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}

func main() {
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))

	fmt.Println(clean2("~/Music/iTunes/../GarageBand/"))

	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(abspath)

	relapth, _ := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/filepath/path.go")
	fmt.Println(relapth)

}
