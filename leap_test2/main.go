package main

import (
	"flag"
	"fmt"

	"github.com/takeweb/golang_lib/date_util"
)

func main() {
	const defaultValue = 1900
	var year int
	flag.IntVar(&year, "year", defaultValue, "year to use")
	flag.Parse()
	fmt.Println(year)
	var result = date_util.IsLeapYear(year)
	fmt.Println(result)
}
