package main

import (
	"fmt"
)

func main() {
	a := []int{10, 20, 30, 40, 50}
	fmt.Println(a)

	a = push(a, 1000)
	fmt.Println(a)

	a = pop(a, 1000)
	fmt.Println(a)

	a = unshift(a, 1000)
	fmt.Println(a)

	a = shift(a, 1000)
	fmt.Println(a)

	a = insert(a, 1000, 2)
	fmt.Println(a)

	a = remove(a, 2)
	fmt.Println(a)
}

// 最後に追加
func push(a []int, v int) []int {
	return append(a, v)
}

// 最後を削除
func pop(a []int, v int) []int {
	return a[:len(a)-1]
}

// 最初に追加
func unshift(a []int, v int) []int {
	return append([]int{v}, a...)
}

// 最初を削除
func shift(a []int, v int) []int {
	return a[1:]
}

// 指定位置に追加
func insert(a []int, v int, p int) []int {
	a = append(a, 0)
	a = append(a[:p+1], a[p:len(a)-1]...)
	a[p] = v
	return a
}

// 指定位置を削除
func remove(a []int, p int) []int {
	return append(a[:p], a[p+1:]...)
}
