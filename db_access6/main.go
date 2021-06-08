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

	ids := console.Input("Input update ID")
	id, _ := strconv.Atoi(ids)

	qry := "SELECT * FROM mydata WHERE id = ?"
	rw := con.QueryRow(qry, id)
	tgt := mydatafmRw(rw)
	fmt.Println(tgt.Str())

	f := console.Input("Delete it? (y/n)")
	if f == "y" {
		qry = "DELETE FROM mydata WHERE id = ?"
		con.Exec(qry, id)
	}

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

// get Mydata from Row.
func mydatafmRw(rs *sql.Row) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Email, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}
