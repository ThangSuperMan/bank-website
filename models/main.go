package models

import (
    "fmt"
	"database/sql"
    "bank/helper")

func InitModel(db *sql.DB) {
    helper.PrintMessages("InitModel")
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

func SetupDB() *sql.DB {
    helper.PrintMessages("setupDB")
	db, err := sql.Open("sqlite3", "./bank_db.db")
    helper.CheckErr(err)
	fmt.Println("Successfully connected!")

	return db
}
