package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 乱数の種を初期化
	rand.Seed(time.Now().UnixNano())

	// コマンドライン引数の数をカウント
	// 設定されていない場合、エラーとする
	c := len(os.Args) - 1
	if c < 1 {
		fmt.Fprintf(os.Stderr, "[useage] %s choice1 choice2...\n", os.Args[0])
		os.Exit(1)
	}

	// 乱数で選択したものを表示
	fmt.Printf("%s\n", os.Args[rand.Intn(c)+1])
}
