package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Data is interface for Mydata.
type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

// Mydata is structure.
type Mydata struct {
	Name string
	Data []int
}

// SetValue is Mydata's method.
func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	md.Data = vali
}

// PrintData is Mydata's method.
func (md *Mydata) PrintData() {
	if md != nil {
		fmt.Println("Name: ", md.Name)
		fmt.Println("Data: ", md.Data)
	} else {
		fmt.Println("**THis is Nil value.**")
	}
}

// Yourdata is structure.
type Yourdata struct {
	Name string
	Mail string
	Age  int
}

// SetValue is Yourdata's method.
func (md *Yourdata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	md.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	md.Age = n
}

// PrintData is Yourdata's method.
func (md *Yourdata) PrintData() {
	fmt.Printf("I'm %s. (%d).\n", md.Name, md.Age)
	fmt.Printf("mail: %s.\n", md.Mail)
}

func main() {
	var obn *Mydata
	obn.PrintData()

	ob := [3]Data{}

	ob[0] = new(Mydata)
	ob[0].SetValue(map[string]string{
		"name": "Sachiko",
		"data": "55, 66, 77",
	})

	ob[1] = new(Yourdata)
	ob[1].SetValue(map[string]string{
		"name": "Maimi",
		"mail": "mami@mume.mo",
		"age":  "34",
	})

	ob[2] = new(Mydata)
	ob[2].SetValue(map[string]string{
		"name": "Taro",
		"data": "45, 56, 87",
	})

	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}
}
