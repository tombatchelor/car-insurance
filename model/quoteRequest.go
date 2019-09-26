package model

import (
	"math/rand"
	"time"
)

type QuoteRequest struct {
	Price int `json:"price`
	Year int `json:"year"`
	Manufacturer string `json:"manufacturer"`
	Excess int `json:"excess"`
	Liability int `json:"liability"`
}

func (quoteRequest *QuoteRequest) Quote() (quoteResponse QuoteResponse) {
	quote := float64(quoteRequest.Price) * 0.05
	quoteResponse.Price = quote
	random := rand.Int31n(5)
	switch random {
	case 0:
		quoteResponse.Company = "Calamity Insurace"
		Sleep(5000)
	case 1:
		quoteResponse.Company = "Accidents R US"
		Sleep(500)
	case 2:
		quoteResponse.Company = "The Lieutenant"
		Sleep(750)
	case 3:
		quoteResponse.Company = "Crashing Pumpkins"
		Sleep(1000)
	case 4:
		quoteResponse.Company = "NeverPayOut Insurance"
		Sleep(600)
	}
	return quoteResponse
}

func Sleep(delay int32) {
	delay += rand.Int31n(750)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
