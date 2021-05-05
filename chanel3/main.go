package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string)
	go first(10, c)
	second(10, c)
	fmt.Println()
}

func prmsg(n int, s string) {
	fmt.Println(s)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func first(n int, c chan string) {
	const nm string = "first-"
	for i := 0; i < 10; i++ {
		s := nm + strconv.Itoa(i)
		prmsg(n, s)
		c <- s
	}
}

func second(n int, c chan string) {
	for i := 0; i < 10; i++ {
		prmsg(n, "second:["+<-c+"]")
	}
}
