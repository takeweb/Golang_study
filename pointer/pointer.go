package main

import (
	"fmt"
)

func main() {
	// アドレスを取得する変数の宣言
	var v int32 = 100

	// アドレス演算子
	var pointer *int32 = &v // ポインタ型の変数にアドレス値を設定
	fmt.Println("ポインタのアドレス値：", pointer)

	// 間接演算子
	fmt.Println("間接演算子の結果：", *pointer)

	// 変数の値を直接更新
	v = 200
	fmt.Println("変数vの更新：", *pointer)

	// ポインターから値を更新
	*pointer = 300
	fmt.Println("ポインタから値を更新：", v)
	fmt.Println("ポインタのアドレス値：", pointer)
}
