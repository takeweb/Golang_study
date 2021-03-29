package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := input("type a price")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
	}
	p := float64(n)
	fmt.Println(int(p * 1.1))
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
