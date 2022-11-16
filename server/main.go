package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func setupDB() *sql.DB {
	connectInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", connectInfo)
	checkErr(err)

	fmt.Println("connectInfo: ", connectInfo)
	fmt.Println("Successfully connected!")
	return db
}

func printMessages(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func main() {
	fmt.Println("hello")
	db := setupDB()
	fmt.Println("db: ", db)
}
