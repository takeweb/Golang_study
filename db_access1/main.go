package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
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

	q := "SELECT * FROM mydata"
	rs, er := con.Query(q)
	if er != nil {
		panic(er)
	}

	for rs.Next() {
		var md Mydata
		er := rs.Scan(&md.ID, &md.Name, &md.Email, &md.Age)
		if er != nil {
			panic(er)
		}
		fmt.Println(md.Str())
	}
}
