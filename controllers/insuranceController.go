package controllers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"car-insurance/model"
	u "car-insurance/utils"
)

var SimpleQuote = func(w http.ResponseWriter, r *http.Request) {
	quoteRequest := &model.QuoteRequest{}

	err := json.NewDecoder(r.Body).Decode(quoteRequest)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		fmt.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quoteRequest.Quote())
}
