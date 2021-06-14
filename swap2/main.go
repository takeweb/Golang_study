package main

import (
	"fmt"
)

func main() {
	a, b := 3, 4
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)
}

func swap2(a *int, b *int) {
	*a, *b = *b, *a
}
