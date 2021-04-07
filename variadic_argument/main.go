package main

import (
	"fmt"
)

func main() {
	m := []string{
		"one", "two", "three",
	}
	fmt.Println(m)

	m = push(m, "1", "2", "3")
	fmt.Println(m)
}

func push(a []string, v ...string) (s []string) {
	s = append(a, v...)
	return
}
