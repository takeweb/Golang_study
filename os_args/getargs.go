package main

import (
	"fmt"
	"os"
)

func main() {
	/* os.Argsの要素数を表示 */
	fmt.Printf("length=%d\n", len(os.Args))

	/* os.Argsの内容を全て表示する */
	for i, v := range os.Args {
		fmt.Println(i, ":", v)
	}
}
