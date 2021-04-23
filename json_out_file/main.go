package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonData struct {
	A       string
	B       string   `json:"code"`
	C       int      `json:",string"`
	D       string   `json:",omitempty"`
	E       string   `json:"-"`
	Targets []target `json:"target"`
}

type target struct {
	Name      string `json:"name"`
	Threshold int    `json:"threshold"`
}

func main() {
	outJsonFile("write.txt")
}

// 構造体をJSONデータに変換してファイルへ出力
func outJsonFile(filename string) {
	data := JsonData{
		A: "データ１",
		B: "データ２",
		C: 2,
		D: "",
		E: "データ５",
		Targets: []target{
			{
				Name:      "taro",
				Threshold: 3},
			{
				Name:      "jiro",
				Threshold: 4}}}
	j, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err_w := ioutil.WriteFile(filename, j, 0666)
	if err_w != nil {
		fmt.Println(err_w)
		os.Exit(1)
	}
}
