package main

import (
	"log"
	"net/http"
	"text/template"
)

// Temps is template structure.
type Temps struct {
	notemp *template.Template
	indx   *template.Template
	helo   *template.Template
}

var flg bool = true

// Template for no-template.
func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

// setup template function.
func setupTemp() *Temps {
	temps := new(Temps)
	temps.notemp = notemp()

	// set index template.
	indx, er := template.ParseFiles("templates/index.html")
	if er != nil {
		indx = temps.notemp
	}
	temps.indx = indx

	// set hello template.
	helo, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		helo = temps.notemp
	}
	temps.helo = helo

	return temps
}

// handler.
func index(w http.ResponseWriter, r *http.Request, tmp *template.Template) {
	er := tmp.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

// hello handler.
func hello(w http.ResponseWriter, r *http.Request, tmp *template.Template) {
	item := struct {
		Flg      bool
		Title    string
		Message  string
		JMessage string
		Items    []string
	}{
		Flg:      flg,
		Title:    "Send Value",
		Message:  "This is Sample message.",
		JMessage: "これはサンプルです。",
		Items:    []string{"One", "Two", "Three"},
	}

	er := tmp.Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
	flg = !flg
}

func main() {
	temps := setupTemp()

	// index handling.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(w, r, temps.indx)
	})

	// index handling.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		hello(w, r, temps.helo)
	})

	http.ListenAndServe(":4000", nil)
}
