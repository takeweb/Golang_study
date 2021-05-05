package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Mydata is json structure.
type Mydata struct {
	Name string
	Mail string
	Tel  string
}

// Str get string value.
func (m *Mydata) Str() string {
	return "<\"" + m.Name + "\" " + m.Mail + ", " + m.Tel + ">"
}

func main() {
	p := "https://tuyano-dummy-data.firebaseio.com/mydata.json"
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}
	defer re.Body.Close()

	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}

	var itms []Mydata
	er = json.Unmarshal(s, &itms)
	if er != nil {
		panic(er)
	}

	for i, im := range itms {
		fmt.Println(i, im.Str())
	}
}
