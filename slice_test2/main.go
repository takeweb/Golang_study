package main

import (
	"fmt"

	"github.com/takeweb/golang_lib/slice_util"
)

func main() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(a)

	a = slice_util.PushInt(a, 1000)
	fmt.Println(a)

	a = slice_util.PopInt(a)
	fmt.Println(a)

	a = slice_util.UnshiftInt(a, 1000)
	fmt.Println(a)

	a = slice_util.ShiftInt(a)
	fmt.Println(a)

	a = slice_util.InsertInt(a, 2, 1000)
	fmt.Println(a)

	a = slice_util.RemoveInt(a, 2)
	fmt.Println(a)
}
