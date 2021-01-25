package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2017, 8, 20, 23, 39, 45, 0, time.Local)
	fmt.Println(date.Format("2006/01/02 15:04:05"))

	mytime := time.Now()
	fmt.Println(mytime.Format("2006/01/02 15:04:05"))
}
