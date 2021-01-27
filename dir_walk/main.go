package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("/etc/", showFullPath)
}

func showFullPath(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}
