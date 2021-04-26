package main

import (
	"fmt"
	"time"
)

func main() {
	go hello("hello", 50)
	hello("bye", 100)
}

func hello(s string, n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("<%d, %s>\n", i, s)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}
