package controllers

import (
	helper "bank/helper"
	"fmt"
	"net/http"
	"text/template"
)

func RenderHomePage(response http.ResponseWriter, request *http.Request) {
    fmt.Println("RouteHome")
    tpl, err := template.ParseGlob("./templates/*")
    helper.CheckErr(err)

    tpl.ExecuteTemplate(response, "main.html", nil)
}
