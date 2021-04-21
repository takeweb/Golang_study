package main

import (
	"encoding/json"
	"io"
	"path/filepath"
	"runtime"

	// "errors"

	// "html/template"
	"log"
	// "net/http"
	"os"
	// "strings"
	// "local.packages/data"
)

type Configuration struct {
	DefDir      string
	DefFilename string
	LogFinename string
}

var config Configuration

var logger *log.Logger

// Convenience function for printing to stdout
// func p(a ...interface{}) {
// 	fmt.Println(a...)
// }

func init() {
	loadConfig()
	file, err := os.OpenFile(config.LogFinename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	// ホームディレクトリを取得
	var configDir string
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		configDir = os.Getenv("APPDATA")
	} else {
		configDir = filepath.Join(home, ".config")
	}

	filename := filepath.Join(configDir, "goinit", "config.json")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// // Convenience function to redirect to the error message page
// func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
// 	url := []string{"/err?msg=", msg}
// 	http.Redirect(writer, request, strings.Join(url, ""), http.StatusFound)
// }

// // Checks if the user is logged in and has a session, if not err is not nil
// func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
// 	cookie, err := request.Cookie("_cookie")
// 	if err == nil {
// 		sess = data.Session{Uuid: cookie.Value}
// 		if ok, _ := sess.Check(); !ok {
// 			err = errors.New("Invalid session")
// 		}
// 	}
// 	return
// }

// // parse HTML templates
// // pass in a list of file names, and get a template
// func parseTemplateFiles(filenames ...string) (t *template.Template) {
// 	var files []string
// 	t = template.New("layout")
// 	for _, file := range filenames {
// 		files = append(files, fmt.Sprintf("templates/%s.html", file))
// 	}
// 	t = template.Must(t.ParseFiles(files...))
// 	return
// }

// func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
// 	var files []string
// 	for _, file := range filenames {
// 		files = append(files, fmt.Sprintf("templates/%s.html", file))
// 	}

// 	templates := template.Must(template.ParseFiles(files...))
// 	templates.ExecuteTemplate(writer, "layout", data)
// }

// for logging
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// version
func version() string {
	return "0.1"
}

// ファイルをコピー
func CopyFile(fromFile string, toFile string) {
	f, err := os.Open(fromFile)
	if err != nil {
		log.Fatal(err)
	}

	t, err := os.Create(toFile)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(t, f)
	if err != nil {
		log.Fatal(err)
	}
}

// ホームディレクトリを取得
func GetHomeDir() string {
	var homeDir string
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		homeDir = os.Getenv("HOMEPATH")
	}
	return homeDir
}
