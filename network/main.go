package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	p := "https://golang.org"
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}
	defer re.Body.Close()

	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}

	fmt.Println(string(s))
}
