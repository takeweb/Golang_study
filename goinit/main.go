package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var dir string
	var module string
	var file string

	// コマンドライン引数取得
	flag.StringVar(&module, "m", "test", "mosule to make")
	flag.StringVar(&dir, "d", filepath.Join(GetHomeDir(), config.DefDir), "base dir to make")
	flag.StringVar(&file, "f", config.DefFilename, "filename to make")
	flag.Parse()

	// ディレクトリ作成
	targetDir := filepath.Join(dir, module)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(module)

	// テンプレートファイルコピー
	fromFile := filepath.Join(dir, "goinit", "template", config.DefFilename)
	toFile := filepath.Join(targetDir, file)
	CopyFile(fromFile, toFile)

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
