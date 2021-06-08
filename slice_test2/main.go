package main

import (
	"fmt"

	"github.com/takeweb/slice"
)

func main() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(a)

	a = slice.PushInt(a, 1000)
	fmt.Println(a)

	a = slice.PopInt(a)
	fmt.Println(a)

	a = slice.UnshiftInt(a, 1000)
	fmt.Println(a)

	a = slice.ShiftInt(a)
	fmt.Println(a)

	a = slice.InsertInt(a, 2, 1000)
	fmt.Println(a)

	a = slice.RemoveInt(a, 2)
	fmt.Println(a)
}
