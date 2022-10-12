package controllers

import (
	"github.com/maickmachado/upvote-api/models"
	"html/template"
	"log"
	"net/http"
)

func ErrorHandler404(w http.ResponseWriter, r *http.Request) {

	TmplError, _ := template.ParseFiles("./template/layout-erro.html")

	w.WriteHeader(http.StatusNotFound)
	data := models.DetailPageData{
		PageTitle: "Erro 404 - Not Found",
	}
	err := TmplError.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func ErrorHandler500(w http.ResponseWriter, r *http.Request) {

	TmplError, _ := template.ParseFiles("./template/layout-erro.html")

	w.WriteHeader(http.StatusInternalServerError)
	data := models.DetailPageData{
		PageTitle: "Erro 500 - Internal Server Error",
	}
	err := TmplError.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
