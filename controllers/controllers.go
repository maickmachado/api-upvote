package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maickmachado/upvote-api/database"
	"github.com/maickmachado/upvote-api/models"
	"html/template"
	"log"
	"net/http"
)

var tmplMainPage = template.Must(template.ParseFiles("template/layout-main-page.html"))
var tmplDetailPage = template.Must(template.ParseFiles("template/layout-detail-page.html"))
var tmpl404 = template.Must(template.ParseFiles("template/layout-404.html"))
var tmplRanking = template.Must(template.ParseFiles("template/layout-ranking.html"))

func GetAllData(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusOK)

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		//Handle Error
	}

	req.Header = http.Header{
		"X-CMC_PRO_API_KEY": {"ab8d13c8-9dae-417c-96a8-02e2e587563d"},
		//"Content-Type":      {"application/json"},
	}

	res, err := client.Do(req)
	if err != nil {
		//Handle Error
	}

	var responseObject models.Response

	json.NewDecoder(res.Body).Decode(&responseObject)
	//
	comp := len(responseObject.CryptoData)

	data := models.MainPageData{
		PageTitle:   "Crypto List",
		CryptoCount: comp,
		Cryptos:     &responseObject,
	}
	tmplMainPage.Execute(w, data)

}

func GetRanking(w http.ResponseWriter, r *http.Request) {

	response, err := database.OrderByVotes()
	if err != nil {

	}

	data := models.PageData{
		PageTitle: "Crypto List",
		//CryptoCount: int(comp),
		CryptosDataBase: response,
	}
	tmplRanking.Execute(w, data)

}

func CryptoDetail(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	var detailResponseObject models.CryptoData
	detailResponseObject.Slug = vars["name"]

	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)

	req.Header = http.Header{
		"X-CMC_PRO_API_KEY": {"ab8d13c8-9dae-417c-96a8-02e2e587563d"},
		//INCLUIR A ZORRA DO COD 200 OU 404 OU ETC AQUI NO HEADER
	}

	res, _ := client.Do(req)

	var responseObject models.Response
	//PORQUE SEM O & NÃO PUXA NENHUM VALOR?
	json.NewDecoder(res.Body).Decode(&responseObject)

	cryptoMatchData, existTrue := database.CheckIfExist(responseObject, detailResponseObject)

	if !existTrue {
		ErrorHandler404(w, r)
	} else {
		singleCrypto, _ := database.GetPending(detailResponseObject.Slug)
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
		tmplDetailPage.Execute(w, data)
	}
}

func VoteCrypto(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	crypto := vars["text"]
	database.Upvote(crypto)
	GetAllData(w, r)
	//w.Header().Set("Content-Type", "application/json")
	//cryptos, _ := database.GetAll()
	////for i, v := range tasks {
	////	w.Write([]byte(fmt.Sprintf("%d: %s - total de votos: %v\n", i+1, v.Text, v.Upvote)))
	////}
	//data := models.PageData{
	//	PageTitle: "Crypto List",
	//	CryptosDataBase:   cryptos,
	//}
	//tmplMainPage.Execute(w, data)
	////w.WriteHeader(http.StatusOK)
	////w.Write([]byte(fmt.Sprintf("TEste")))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Some Error Occurred"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
