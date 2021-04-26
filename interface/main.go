package main

import (
	"fmt"
)

// Data is interface.
type Data interface {
	Initialize(name string, data []int)
	PrintData()
}

type Mydata struct {
	Name string
	Data []int
}

func main() {
	var ob Mydata = Mydata{}
	// var ob Data = new(Mydata)
	ob.Initialize("Sachiko", []int{56, 66, 77})
	ob.PrintData()
	ob.Check()
}

// Initialize is init method.
func (md *Mydata) Initialize(name string, data []int) {
	md.Name = name
	md.Data = data
}

// PrintDAta is println all data.
func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

// Check is method
func (md *Mydata) Check() {
	fmt.Printf("Check! [%s]\n", md.Name)
}
