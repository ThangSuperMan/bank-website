package main

import (
	"bank/controllers"
	"bank/helper"
	"bank/models"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
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

func handlers(router *mux.Router) {
	helper.PrintMessages("handlers")
	router.HandleFunc("/", controllers.RenderHomePage).Methods("get")
	router.HandleFunc("/user/register", controllers.RenderRegisterPage).Methods("get")
}

func main() {
	fmt.Println("hello")
	db := models.SetupDB()
	models.InitModel(db)
	router := mux.NewRouter()

	// Serve static filss
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(s)
	handlers(router)
	http.ListenAndServe(":3002", router)
}
