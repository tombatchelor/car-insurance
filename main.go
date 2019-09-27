package main

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
	"car-insurance/controllers"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/metrics", promhttp.Handler())

	router.HandleFunc("/insurance/simpleQuote",  controllers.SimpleQuote).Methods("POST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
