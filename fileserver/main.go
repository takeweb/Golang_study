package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":4000", http.FileServer(http.Dir(".")))
}
