package main

import (
	"fmt"

	"github.com/takeweb/slice"
)

func main() {
	var a interface{} = []int{10, 20, 30, 40, 50}
	fmt.Println(a)

	slice.Push(&a, 1000)
	fmt.Println(a)

	slice.Pop(&a)
	fmt.Println(a)

	slice.Unshift(&a, 1000)
	fmt.Println(a)

	slice.Shift(&a)
	fmt.Println(a)

	slice.Insert(&a, 2, 1000)
	fmt.Println(a)

	slice.Remove(&a, 2)
	fmt.Println(a)
}
