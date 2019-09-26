package model

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
	quoteResponse.Company = "Calamity Insurace"
	return quoteResponse
}
