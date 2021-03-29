package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := input("type your name")
	fmt.Println("Hello, " + name + "!")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
