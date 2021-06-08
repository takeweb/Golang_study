package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/takeweb/console"
)

// Mydata is structure
type Mydata struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// Str get string value.
func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\" " + m.Email + "," + strconv.Itoa(m.Age) + ">"
}

var qry string = "SELECT * FROM mydata WHERE id = ?"

func main() {
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	for {
		s := console.Input("id")
		if s == "" {
			break
		}

		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}

		rs := con.QueryRow(qry, n)
		var md Mydata
		er2 := rs.Scan(&md.ID, &md.Name, &md.Email, &md.Age)
		if er2 != nil {
			panic(er)
		}
		fmt.Println(md.Str())
	}
	fmt.Println("***end***")
}
