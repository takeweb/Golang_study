package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go total(c)
	c <- 100
	time.Sleep(100 * time.Millisecond)
}

func total(c chan int) {
	n := <-c
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
}
