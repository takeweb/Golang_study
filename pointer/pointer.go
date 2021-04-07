package main

import (
	"fmt"
)

func main() {
	// アドレスを取得する変数の宣言
	// var v int32 = 100
	v := 100
	fmt.Println("元の変数vの値：", v)

	// アドレス演算子
	// var p *int32 = &v
	p := &v // ポインタ型の変数にアドレス値を設定
	fmt.Println("ポインタのアドレス値：", p)

	// 間接演算子
	fmt.Println("間接演算子の結果：", *p)

	// 変数の値を直接更新
	v = 200
	fmt.Println("変数vの更新：", *p)

	// ポインターから値を更新
	*p = 300
	fmt.Println("ポインタから値を更新：", v)
	fmt.Println("ポインタのアドレス値：", p)
}
