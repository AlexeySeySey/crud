package servs

import (
	_ "../sql/go-sql-driver/mysql"
	"net/http"
)

type Note struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Text    string    `json:"text"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

const views = "./templates"
/*
func (note *Note) Main(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("../templates/index.html")
	view.Execute(w, "1")
}*/

func (note *Note) FetchAll(w http.ResponseWriter, r *http.Request) {

	/*db, err := sql.Open("mysql", "root:16d08d2000d@tcp(127.0.0.1:3306)/GoSome")
	simple.Err(err, func() {
		log.Fatal(err)
	})

	query, err := db.Query("SELECT * FROM Notes")

	simple.Err(err,func(){
		fmt.Fprint(w,err)
	})

	for query.Next() {

         err := query.Scan(&note.Id, &note.Name, &note.Text, &note.Created, &note.Updated)
         simple.Err(err,func(){
         	log.Fatal(err.Error())
		 })

	}

	view, _ := template.ParseFiles("./templates/test.html")
	view.Execute(w,note)

	query.Close()
	db.Close()*/

}
