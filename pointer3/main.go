package main

import (
	"fmt"
)

func main() {
	n := 123
	p1 := &n
	fmt.Printf("p1 value:%d, address:%p\n", *p1, p1)
	q1 := &p1
	fmt.Printf("q1 value:%d, address:%p\n", **q1, q1)

	m := 1000
	p2 := &m
	fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)
	q2 := &p2
	fmt.Printf("q2 value:%d, address:%p\n", **q2, q2)

	pb := p1
	p1 = p2
	p2 = pb

	fmt.Printf("p1 value:%d, address:%p\n", *p1, p1)
	fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)
	fmt.Printf("q1 value:%d, address:%p\n", **q1, *q1)
	fmt.Printf("q2 value:%d, address:%p\n", **q2, *q2)
}
