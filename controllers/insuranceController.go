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

	resp := u.Message(true, "success")
	resp["quoteResponse"] = quoteRequest.Quote()
	u.Respond(w, resp)
}
