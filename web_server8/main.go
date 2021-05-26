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

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

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
	msg := "login name and password:"
	ses, _ := cs.Get(r, "hello-session")

	if r.Method == "POST" {
		ses.Values["login"] = nil
		ses.Values["name"] = nil
		nm := r.FormValue("name")
		pw := r.FormValue("pass")
		if nm == pw {
			ses.Values["login"] = true
			ses.Values["name"] = nm
		}
		ses.Save(r, w)
	}

	flg, _ := ses.Values["login"].(bool)
	lname, _ := ses.Values["name"].(string)
	if flg {
		msg = "logined: " + lname
	}

	item := struct {
		Title   string
		Message string
	}{
		Title:   "Send Values",
		Message: msg,
	}

	er := tmp.Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
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
