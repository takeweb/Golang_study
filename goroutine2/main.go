package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	msg := "start"
	prmsg := func(nm string, n int) {
		fmt.Println(nm, msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	hello := func(n int) {
		const nm string = "hello"
		for i := 0; i < 10; i++ {
			msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
	}

	main := func(n int) {
		const nm string = "*main"
		for i := 0; i < 5; i++ {
			msg += " m" + strconv.Itoa(i)
			prmsg(nm, 100)
		}
	}

	go hello(60)
	main(100)
}
