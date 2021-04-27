package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go total(1000, c)
	go total(100, c)
	go total(10, c)
	// fmt.Println("total:", <-c)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z)
}

func total(n int, c chan int) {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	c <- t
}
