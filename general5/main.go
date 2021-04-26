package main

import (
	"fmt"
	"reflect"
)

// General is all type data.
type General interface{}

// GData is holding Genaral value.
type GData interface {
	Set(nm string, g General) GData
	Print()
}

// NData is structure.
type NData struct {
	Name string
	Data []int
}

// Set is NData method.
func (nd *NData) Set(nm string, g General) GData {
	nd.Name = nm
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf(0)) {
		nd.Data = g.([]int)
	}
	return nd
}

// Print is NData method.
func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

// SData is structure.
type SData struct {
	Name string
	Data []string
}

// Set is SData method.
func (sd *SData) Set(nm string, g General) GData {
	sd.Name = nm
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf("")) {
		sd.Data = g.([]string)
	}
	return sd
}

// Print is SData method.
func (sd *SData) Print() {
	fmt.Printf("* %s [%s] * \n", sd.Name, sd.Data)
}

func main() {
	var data = []GData{}
	data = append(data, new(NData).Set("Taro", []int{1, 2, 3}))
	data = append(data, new(SData).Set("Jiro", []string{"hello!", "bye"}))
	data = append(data, new(NData).Set("Hanako", "98700"))
	data = append(data, new(SData).Set("Sachiko", "happy?"))
	for _, ob := range data {
		ob.Print()
	}
}
