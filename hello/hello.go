package main

import (
	"fmt"
	"net/http"
)

// ServerHTTPメソッド用の構造体型
type Server struct{}

// httpリクエストを受け取るメソッド
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// HTML文字列
	h := `
	<html>
		<head><title>Hello</title></head>
		<body>
			Hello
		</body>
	</html>
	`
	// ブラウザにHTMLを送信
	fmt.Fprint(w, h)
}

func main() {
	// Webサーバを起動
	http.ListenAndServe(":4000", Server{})
}

