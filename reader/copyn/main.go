package main

import (
	"io"
	"os"
	"strings"
)

func copyN(dest io.Writer, src io.Reader, length int) (n int64, err error) {
	written, err := io.Copy(dest, io.LimitReader(src, n))
	if err != nil {
		return 0, err
	}
	return written, nil
}

func main() {
	reader := strings.NewReader("1234567890")
	copyN(os.Stdout, reader, 5)
}
