package main

import (
	"os"
)

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/Users/dengrunting/Desktop/test.txt")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break

		}
		os.Stdout.Write(buf[:n])
	}
}
