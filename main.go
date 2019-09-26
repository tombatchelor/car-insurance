package main

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
	"car-insurance/controllers"
)

func main() {
	router := mux.NewRouter()

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
