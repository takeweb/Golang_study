package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func printIf(b bool) {
	var funcVar func(int, int) int

	if b {
		funcVar = add
	} else {
		funcVar = mul
	}

	fmt.Println(funcVar(3, 4))
}

func main() {
	printIf(true)
	printIf(false)
}
