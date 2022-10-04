package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PageData struct {
	PageTitle       string
	CryptosDataBase []*CryptoDataBase
}

type CryptoDataBase struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Name      string             `json:"name" bson:"name"`
	Upvote    int                `json:"votes" bson:"votes"`
}

type MainPageData struct {
	PageTitle   string
	CryptoCount int
	Cryptos     *Response
}

type DetailPageData struct {
	PageTitle string
	Cryptos   *CryptoData
	Votes     int
}

type Response struct {
	CryptoData []CryptoData `json:"data"`
}

type CryptoData struct {
	ID      int    `json:"id"`
	CmcRank int    `json:"cmc_rank"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Slug    string `json:"slug"`
	Quote   Quote
}

type Quote struct {
	Usd Usd `json:"USD"`
}

type Usd struct {
	Price float64 `json:"price"`
}
