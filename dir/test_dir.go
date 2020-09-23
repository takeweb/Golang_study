package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	paths := dirWalk("./test")
	for _, path := range paths {
		fmt.Println(path)
	}
}

func dirWalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirWalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths
}
