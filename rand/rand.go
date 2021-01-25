package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// var num = rand.Intn(10) + 1
	var num = rand.Intn(11)
	fmt.Println(num)
	//num = rand.Intn(10) + 1
	num = rand.Intn(11)
	fmt.Println(num)
}
