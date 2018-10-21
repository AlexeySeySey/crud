package servs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "../sql/go-sql-driver/mysql"
)

const (
	views = "./templates"
)

var DB = map[string]string{
	"RDBMS":    "mysql",
	"NAME":     "root",
	"PASSWORD": "16d08d2000d",
	"ADDRESS":  "127.0.0.1",
	"PORT":     "3306",
	"DBNAME":   "GoSome",
	"PROTOCOL": "tcp",
}

func inits() *sql.DB {
	db, err := sql.Open(DB["RDBMS"], DB["NAME"]+":"+DB["PASSWORD"]+"@"+DB["PROTOCOL"]+"("+DB["ADDRESS"]+":"+DB["PORT"]+")/"+DB["DBNAME"])
	if err != nil {
		log.Fatal(err)
	}
	return db
}

var db = inits()

type Note struct {
	Id      int
	Name    string
	Text    string
	Updated string
	Created string
}

func Main(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, views+"/index.html")
}

type Data struct {
	Tabledata []*Note
}

func (note *Note) FetchAll(w http.ResponseWriter, r *http.Request) {

	d := Data{}
	rows, err := db.Query("SELECT * FROM Notes")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	rs := make([]*Note, 0)
	for rows.Next() {
		rst := new(Note)
		err := rows.Scan(&rst.Id, &rst.Name, &rst.Text, &rst.Updated, &rst.Created)
		if err != nil {
			log.Println(err)
		}

		rs = append(rs, rst)
	}

	d.Tabledata = rs
	json.NewEncoder(w).Encode(d.Tabledata)
}

func (note *Note) NewOne(w http.ResponseWriter, r *http.Request) {

	// if created != nil -> change only updated + created err

	var name string = r.FormValue("name")
	var text string = r.FormValue("text")

	in, err := db.Prepare("insert into Notes (name,text,updated) values(?,?,NOW())")
	if err != nil {
		log.Fatal(err)
	}
	in.Exec(name, text)

	fmt.Fprint(w, "Added")
}
