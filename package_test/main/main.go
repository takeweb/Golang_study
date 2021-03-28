package main

import (
	"workpkg/workhello"

	"local.packages/workhello2"
	"local.packages/workhello3"
)

func main() {
	workhello.Hello()
	workhello2.Hello()
	workhello3.Hello()
}
