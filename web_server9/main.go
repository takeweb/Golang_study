package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

// Temps is template structure.
type Temps struct {
	notemp *template.Template
	indx   *template.Template
	helo   *template.Template
}

// session variable. (not used)
var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

// Template for no-template.
func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

// get target Template.
func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("templates/"+fname+".html", "templates/head.html", "templates/foot.html")
	return tmps
}

// index handler.
func index(w http.ResponseWriter, r *http.Request) {
	item := struct {
		Template string
		Title    string
		Message  string
	}{
		Template: "index",
		Title:    "Index",
		Message:  "This is Top page.",
	}

	er := page(item.Template).Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// hello handler.
func hello(w http.ResponseWriter, r *http.Request) {
	data := []string{
		"One", "Two", "Three",
	}

	item := struct {
		Template string
		Title    string
		Data     []string
	}{
		Template: "hello",
		Title:    "Hello",
		Data:     data,
	}

	er := page(item.Template).Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// main program.
func main() {
	// index handling.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(w, r)
	})

	// index handling.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		hello(w, r)
	})

	http.ListenAndServe(":4000", nil)
}
