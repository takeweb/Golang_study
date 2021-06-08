package main

import (
	"flag"
	"fmt"

	"github.com/takeweb/date"
)

func main() {
	const defaultValue = 1900
	var year int
	flag.IntVar(&year, "year", defaultValue, "year to use")
	flag.Parse()
	fmt.Println(year)
	var result = date.IsLeapYear(year)
	fmt.Println(result)
}
