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

var qry string = "SELECT * FROM mydata WHERE name LIKE ? OR mail LIKE ?"

func main() {
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	for {
		s := console.Input("find from Name or Email")
		if s == "" {
			break
		}

		rs, er := con.Query(qry, "%"+s+"%", "%"+s+"%")
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
	fmt.Println("***end***")
}
