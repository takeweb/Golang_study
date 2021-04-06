package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("name >")
	fmt.Scan(&name)
	fmt.Println("Hello, " + name + "!")
}
