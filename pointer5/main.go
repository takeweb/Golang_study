package main

import (
	"fmt"
)

func main() {
	ar := []int{10, 20, 30}
	fmt.Println(ar)
	initialize(&ar)
	fmt.Println(ar)
}

func initialize(ar *[]int) {
	for i := 0; i < len(*ar); i++ {
		(*ar)[i] = 0
	}
}
