package main

import (
	"fmt"
)

func main() {
	n := 123
	p1 := &n
	fmt.Printf("p1 value:%d, address:%p\n", *p1, p1)

	m := 1000
	p2 := &m
	fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)

	pb := p1
	p1 = p2
	p2 = pb

	fmt.Printf("p1 value:%d, address:%p\n", *p1, p1)
	fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)
}
