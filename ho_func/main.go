package main

import "fmt"

func main() {
	modify := func(a []string, f func([]string) []string) []string {
		return f(a)
	}

	m := []string{
		"1st", "2nd", "3rd",
	}
	fmt.Println(m)

	m1 := modify(m, func([]string) []string {
		return append(m, m...)
	})
	fmt.Println(m1)

	m2 := modify(m, func([]string) []string {
		return m[:len(m)-1]
	})
	fmt.Println(m2)

	m3 := modify(m, func([]string) []string {
		return m[1:]
	})
	fmt.Println(m3)
}
