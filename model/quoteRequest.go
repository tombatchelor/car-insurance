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
	currentTime := time.Now()
	var random int32
	if currentTime.Hour()%2 == 0 {
		random = rand.Int31n(4)
	} else {
		random = rand.Int31n(5)
	}
	switch random {
	case 0:
		quoteResponse.Company = "NeverPayOut Insurance"
		Sleep(600)
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
		quoteResponse.Company = "Calamity Insurance"
		Sleep(10000)
	}
	return quoteResponse
}

func Sleep(delay int32) {
	delay += rand.Int31n(750)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
