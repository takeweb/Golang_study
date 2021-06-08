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
	ae := strconv.Itoa(tgt.Age)
	nm := console.Input("name(" + tgt.Name + ")")
	ml := console.Input("mail(" + tgt.Email + ")")
	ge := console.Input("name(" + ae + ")")
	ag, _ := strconv.Atoi(ge)

	if nm == "" {
		nm = tgt.Name
	}

	if ml == "" {
		ml = tgt.Email
	}

	if ge == "" {
		ag = tgt.Age
	}

	qry = "UPDATE mydata SET name = ?, mail = ?, age = ? WHERE id = ?"
	con.Exec(qry, nm, ml, ag, id)

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
