package main

import (
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "you can also add header")
	request.Write(os.Stdout)

}
