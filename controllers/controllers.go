package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maickmachado/upvote-api/database"
	"github.com/maickmachado/upvote-api/models"
	"html/template"
	"io"
	"log"
	"net/http"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	TmplMainPage, _ := template.ParseFiles("./template/layout-main-page.html")

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		ErrorHandler500(w, r)
	} else {
		req.Header = http.Header{
			"X-CMC_PRO_API_KEY": {"ab8d13c8-9dae-417c-96a8-02e2e587563d"},
		}

		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}

		var responseObject models.Response

		err = json.NewDecoder(res.Body).Decode(&responseObject)
		if err != nil {
			log.Println(err)
		}
		//
		comp := len(responseObject.CryptoData)

		data := models.MainPageData{
			PageTitle:   "Crypto List",
			CryptoCount: comp,
			Cryptos:     &responseObject,
		}
		err = TmplMainPage.Execute(w, data)
		if err != nil {
			ErrorHandler500(w, r)
		}
	}
}

func GetRanking(w http.ResponseWriter, r *http.Request) {

	TmplRanking, _ := template.ParseFiles("./template/layout-ranking.html")

	response, err := database.OrderByVotes()
	if err != nil {
		log.Println(err)
	}

	data := models.PageData{
		PageTitle:       "Crypto List",
		CryptosDataBase: response,
	}
	err = TmplRanking.Execute(w, data)
	if err != nil {
		ErrorHandler500(w, r)
	}
}

func CryptoDetail(w http.ResponseWriter, r *http.Request) {
	TmplDetailPage, _ := template.ParseFiles("./template/layout-detail-page.html")

	vars := mux.Vars(r)
	var detailResponseObject models.CryptoData
	detailResponseObject.Slug = vars["name"]

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		ErrorHandler500(w, r)
	} else {
		req.Header = http.Header{
			"X-CMC_PRO_API_KEY": {"ab8d13c8-9dae-417c-96a8-02e2e587563d"},
		}

		res, _ := client.Do(req)

		var responseObject models.Response
		err = json.NewDecoder(res.Body).Decode(&responseObject)
		if err != nil {
			log.Println(err)
		}

		cryptoMatchData, existTrue := database.CheckIfExist(responseObject, detailResponseObject)

		if !existTrue {
			ErrorHandler404(w, r)
		} else {
			singleCrypto, err := database.GetCryptoInfoDataBase(detailResponseObject.Slug)
			if err != nil {
				log.Println(err)
			}

			var vote int

			if len(singleCrypto) == 0 {
				vote = 0
			} else {
				vote = singleCrypto[0].Upvote
			}

			data := models.DetailPageData{
				PageTitle: "Crypto List",
				Cryptos:   cryptoMatchData,
				Votes:     vote,
			}
			err = TmplDetailPage.Execute(w, data)
			if err != nil {
				ErrorHandler500(w, r)
			}
		}
	}
}

func VoteCrypto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	crypto := vars["text"]
	database.Upvote(crypto)
	GetAllData(w, r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
