package main

import "fmt"

func main() {
	data := "*新しい値*"
	m1 := modify(data)

	data = "+new data+"
	m2 := modify(data)

	fmt.Println(m1())
	fmt.Println(m2())
}

func modify(d string) func() []string {
	m := []string{
		"1st", "2nd",
	}
	return func() []string {
		return append(m, d)
	}
}
