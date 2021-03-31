package main

import (
	"fmt"

	"github.com/takeweb/golang_lib/slice_util"
)

func main() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(a)

	a = slice_util.Push(a, 1000)
	fmt.Println(a)

	a = slice_util.Pop(a, 1000)
	fmt.Println(a)

	a = slice_util.Unshift(a, 1000)
	fmt.Println(a)

	a = slice_util.Shift(a, 1000)
	fmt.Println(a)

	a = slice_util.Insert(a, 1000, 2)
	fmt.Println(a)

	a = slice_util.Remove(a, 2)
	fmt.Println(a)
}
