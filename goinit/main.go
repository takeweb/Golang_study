package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	const defDir string = "/Users/taketomooishi/dev/Golang_study/"
	const defFilename string = "main.go"

	var dir string
	var module string
	var file string

	flag.StringVar(&module, "m", "test", "mosule to make")
	flag.StringVar(&dir, "d", defDir, "base dir to make")
	flag.StringVar(&file, "f", defFilename, "filename to make")
	flag.Parse()

	// ディレクトリ作成
	targetDir := filepath.Join(dir, module)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(module)

	// テンプレートファイルコピー
	fromFile := filepath.Join(dir, "goinit", "template", defFilename)
	toFile := filepath.Join(targetDir, file)
	copyFile(fromFile, toFile)

	// ターゲットディレクトリへ移動
	if err := os.Chdir(targetDir); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// カレントディレクトリを取得
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(currentDir)

	// 外部コマンド実行
	err_o := exec.Command("go", "mod", "init", module).Run()
	if err_o != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func copyFile(fromFile string, toFile string) {
	f, err := os.Open(fromFile)
	if err != nil {
		log.Fatal(err)
	}

	t, err := os.Create(toFile)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(t, f)
	if err != nil {
		log.Fatal(err)
	}
}
