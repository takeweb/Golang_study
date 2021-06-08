package main

import (
	"fmt"
	"strings"
)

func main() {
	// printIf(120)
	// printIf("bye")
	// printIf([]string{"apple", "orange", "banana"})
	// printIf([2]string{"good", "bye"})
	a := []string{"10", "20", "30", "40", "50"}
	PrintIf(a, "b")
	fmt.Println(Push(a, "b"))
}

func PrintIf(src interface{}, a interface{}) {
	switch value := src.(type) {
	case int:
		fmt.Printf("parameter is integer. [value: %d]\n", value)
	case string:
		value = strings.ToUpper(value) // 対象がstring型なのでstringを引数に取る関数が実行できる
		fmt.Printf("parameter is string. [value: %s]\n", value)
	case []string:
		if add, ok := a.(string); ok {
			value = append(value, add) // 対象がsliceなのでAppendができる
			fmt.Printf("parameter is slice string. [value: %s]\n", value)
		}
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", src)
	}
}

// 最後に追加
func Push(a interface{}, v interface{}) interface{} {
	var rtn interface{}
	switch value := v.(type) {
	case int:
		if array, ok := a.([]int); ok {
			rtn = PushInt(array, value)
		}
	case string:
		if array, ok := a.([]string); ok {
			rtn = PushString(array, value)
		}
	default:
		fmt.Printf("parameter is unknown type. [valueType: %T]\n", v)
	}
	return rtn
}

// 最後に追加
func PushInt(a []int, v ...int) []int {
	return append(a, v...)
}

// 最後に追加
func PushString(a []string, v ...string) []string {
	return append(a, v...)
}
