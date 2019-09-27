package controllers

import (
	"net/http"
	"log"
	"encoding/json"
	"car-insurance/model"
	u "car-insurance/utils"
)

var SimpleQuote = func(w http.ResponseWriter, r *http.Request) {
	quoteRequest := &model.QuoteRequest{}

	err := json.NewDecoder(r.Body).Decode(quoteRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		log.Print(err)
		return
	}
	log.Print(quoteRequest)

	w.Header().Add("Content-Type", "application/json")
	response := quoteRequest.Quote()
	log.Print(response)
	json.NewEncoder(w).Encode(response)
}
