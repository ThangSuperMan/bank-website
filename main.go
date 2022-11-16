package main

import (
	"database/sql"
    helper "bank/helper"
	"fmt"
    "embed"
	"net/http"
	"text/template"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var templateFiles embed.FS

const (
	host     = "localhost"
	port     = 5432
	user     = "thangjenny"
	password = ""
	dbname   = "bank_db"
)

type User struct {
	Username string `json:"username`
	Password string `json:"password`
}

type JsonResponse struct {
	Type    string `json:"type`
	Data    []User `json:"data`
	Message string `json:"message`
}


func setupDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./bank_db.db")
    helper.CheckErr(err)
	fmt.Println("Successfully connected!")

	return db
}

func printMessages(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func Init(db *sql.DB) {
  var stmt string = `
    create table if not exists users (
      id_user integer primary key autoincrement,
      username text not null,
      password text not null,
      profile_name text not null,
      avatar_name text not null
    )
    ` 

  _, err := db.Exec(stmt)
  helper.CheckErr(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("homeHandler")
  // tpl, _ := template.ParseFiles(
  //   "views/hello.html",
  //   "views/common.html",
  //   )
    
    tpl, err := template.ParseGlob("./templates/*")


    helper.CheckErr(err)

    fmt.Println(tpl.DefinedTemplates())
    tpl.ExecuteTemplate(w, "main.html", nil)


    // fmt.Println("Current template is being used: ", tpl.Name())
    // fmt.Println(tpl.DefinedTemplates())
    //
    // tpl.Execute(w, map[string]string {
    //   "name": "youngmoon",
    // })

}

func main() {
    db := setupDB()
    Init(db)
    router := mux.NewRouter()

    s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
    router.PathPrefix("/static/").Handler(s)

    // router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
    router.HandleFunc("/", homeHandler).Methods("get")
    http.ListenAndServe(":3000", router)
}
