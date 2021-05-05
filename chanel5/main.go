package main

import (
	"fmt"
	"time"
)

func main() {
	n1, n2, n3 := 3, 5, 10
	m1, m2, m3 := 100, 75, 50

	c1 := make(chan int)
	go count(n1, m1, c1)

	c2 := make(chan int)
	go count(n2, m2, c2)

	c3 := make(chan int)
	go count(n3, m3, c3)

	for i := 0; i < n1+n2+n3; i++ {
		select {
		case re := <-c1:
			fmt.Println("*  first ", re)
		case re := <-c2:
			fmt.Println("** second", re)
		case re := <-c3:
			fmt.Println("***third ", re)
		}
	}
	fmt.Println("*** finish ***")
}

func count(n int, s int, c chan int) {
	for i := 1; i <= n; i++ {
		c <- i
		time.Sleep(time.Duration(s) * time.Millisecond)
	}
}
