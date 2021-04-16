package main

import (
	"fmt"
)

func main() {
	n := 123

	f := func() {
		fmt.Println(n)
		n++
	}

	f()
	fmt.Println(n)

	n = 20
	f()
	fmt.Println(n)
}
