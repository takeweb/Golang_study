package main

import (
	"fmt"
	"os"
)

func main() {

	// filepath.Join("2021", file.Name())
	if err := os.MkdirAll("hoge/fuga", 0777); err != nil {
		fmt.Println(err)
	}
}
