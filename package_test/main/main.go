package main

import (
	"workpkg/workhello"

	"../package3/workhello3"
	"./package2/workhello2"
)

func main() {
	workhello.Hello()
	workhello2.Hello()
	workhello3.Hello()
}
