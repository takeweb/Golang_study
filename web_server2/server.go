package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// Server ServeHTTPメソッド用構造体
type Server struct{}

// httpリクエストを受け取るメソッド
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// フォームの入力値を取得
	v := r.FormValue("input_value")

	// HTML文字列
	h := `
	<html>
		<head>
			<title>HTMLフォーム</title>
		</head>
		<body>
			<form action="/" method="post">
				<input type="text" name="input_value">
				<input type="submit" name="送信"><br>
				入力値：` + html.EscapeString(v) + `
			</form>
		</body>
	</html>
	`
	// ブラウザにHTMLを送信
	fmt.Fprint(w, h)
}

func main() {
	http.ListenAndServe(":4000", Server{})
	log.Println("start http listening :4000")
}
