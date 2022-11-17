package controllers

import (
	"bank/helper"
	"fmt"
	"net/http"
	"text/template"
)

func RenderRegisterPage(response http.ResponseWriter, request *http.Request) {
    fmt.Println("RenderRegisterPage")
    tpl, err := template.ParseGlob("./templates/*")

    helper.CheckErr(err)

    tpl.ExecuteTemplate(response, "register.html", nil)

}

func HandleRegister(response http.ResponseWriter, request *http.Request) {
    fmt.Println("HandleRegister")
}
