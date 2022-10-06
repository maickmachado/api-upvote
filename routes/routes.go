package routes

import (
	"github.com/maickmachado/upvote-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/cryptocoins", controllers.GetAllData).Methods("GET")
	myRouter.HandleFunc("/ranking", controllers.GetRanking).Methods("GET")
	myRouter.HandleFunc("/cryptocoins/{name}", controllers.CryptoDetail).Methods("GET")
	myRouter.HandleFunc("/healthcheck", controllers.HealthCheck).Methods("GET")
	myRouter.HandleFunc("/cryptocoins/vote/{text}", controllers.VoteCrypto).Methods("POST")
	myRouter.NotFoundHandler = http.Handler(http.HandlerFunc(controllers.ErrorHandler404))

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
