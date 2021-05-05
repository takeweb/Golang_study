package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)
	cr := make(chan int)
	go total(cs, cr)
	cs <- 100
	fmt.Println("total:", <-cr)
}

func total(cs chan int, cr chan int) {
	n := <-cs
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	cr <- t
}
