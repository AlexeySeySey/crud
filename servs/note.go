package servs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	indexPage = "public/index.html"
)

var DB = map[string]string{
	"RDBMS":    "mysql",
	"NAME":     "root",
	"PASSWORD": "root",
	"ADDRESS":  "127.0.0.1",
	"PORT":     "3306",
	"DBNAME":   "GoCrud",
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
	http.ServeFile(w, r, indexPage)
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

func NewOne(w http.ResponseWriter, r *http.Request) {

	var name string = r.FormValue("name")
	var text string = r.FormValue("text")

	in, err := db.Prepare("insert into Notes (name,text,updated) values(?,?,NOW())")
	if err != nil {
		log.Fatal(err)
	}
	in.Exec(name, text)

}

func DropOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	del, err := db.Prepare("delete from Notes where id = ?")
	if err != nil {
		log.Fatal(nil)
	}

	del.Exec(vars["id"])

}

func UpdateOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	fmt.Fprint(w, vars)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	upd, err := db.Prepare("update Notes set name = ?, text = ?, updated = NOW() where id = ?")
	if err != nil {
		log.Fatal(err)
	}

	upd.Exec(vars["name"], vars["text"], id)

}
