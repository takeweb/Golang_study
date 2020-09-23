package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if err := readBytes("foo.txt"); err != nil {
		//fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func readBytes(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 64)
	for {
		c, err := file.Read(buf)
		if c == 0 {
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		line := string(buf[:c])
		fmt.Println(line)
	}
	return nil
}
