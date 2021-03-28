package main

import (
	"flag"
	"fmt"

	"local.packages/util"
)

func main() {
	const defaultValue = 1900
	var year int
	flag.IntVar(&year, "year", defaultValue, "year to use")
	flag.Parse()
	fmt.Println(year)
	var result = util.IsLeapYear(year)
	fmt.Println(result)
}
