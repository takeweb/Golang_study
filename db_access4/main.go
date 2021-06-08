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

func main() {
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	nm := console.Input("Input name")
	nl := console.Input("Input mail")
	age := console.Input("Input age")
	ag, _ := strconv.Atoi(age)

	qry := "INSERT INTO mydata(name, mail, age) values (?, ?, ?)"
	con.Exec(qry, nm, nl, ag)
	showRec(con)
}

// print all record.
func showRec(con *sql.DB) {
	qry := "SELECT * FROM mydata"
	rs, _ := con.Query(qry)
	for rs.Next() {
		fmt.Println(mydatafmRws(rs).Str())
	}
}

// get Mydata from Rows.
func mydatafmRws(rs *sql.Rows) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Email, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}

// // get Mydata from Row.
// func mydatafmRw(rs *sql.Row) *Mydata {
// 	var md Mydata
// 	er := rs.Scan(&md.ID, &md.Name, &md.Email, &md.Age)
// 	if er != nil {
// 		panic(er)
// 	}
// 	return &md
// }
