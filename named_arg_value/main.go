package main

import (
	"fmt"
)

func main() {
	m := []string{
		"one", "two", "three",
	}
	fmt.Println(m)

	m = insert(m, "*", 2)
	fmt.Println(m)

	m = insert(m, "*", 1)
	fmt.Println(m)
}

func insert(a []string, v string, p int) (s []string) {
	s = append(a, "")
	fmt.Println(s)
	s = append(s[:p+1], s[p:len(s)-1]...)
	fmt.Println(s)
	s[p] = v
	return
}
