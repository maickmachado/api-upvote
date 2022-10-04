package controllers

import (
	"github.com/maickmachado/upvote-api/models"
	"net/http"
)

func ErrorHandler404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	data := models.DetailPageData{
		PageTitle: "Erro 404 - Página não encontrada",
	}
	tmpl404.Execute(w, data)
}

func ErrorHandler500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	data := models.DetailPageData{
		PageTitle: "Erro 404 - Página não encontrada",
	}
	tmpl404.Execute(w, data)
}
