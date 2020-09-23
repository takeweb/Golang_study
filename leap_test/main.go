package main

import (
	"flag"
	"fmt"
	"workpkg/dateutil"
)

func main() {
	const defaultValue = 1900
	var year int
	//	var year2 = flag.Int("year", defaultValue, "use")
	flag.IntVar(&year, "year", defaultValue, "year to use")
	flag.Parse()
	//fmt.Println(year)
	var result = dateutil.IsLeap(year)
	fmt.Println(result)

	//	year3 := *year2
	//	result = dateutil.IsLeap(year3)
	//fmt.Println(result)
}
