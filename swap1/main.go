package main

import (
	"fmt"
)

func main() {
	a, b := 3, 4
	fmt.Println(a, b)
	a, b = swap1(a, b)
	fmt.Println(a, b)
}

func swap1(a int, b int) (int, int) {
	return b, a
}
