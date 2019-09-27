package main

import (
	"github.com/gorilla/mux"
	"os"
	"log"
	"net/http"
	"car-insurance/controllers"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/openzipkin/zipkin-go"
    "github.com/openzipkin/zipkin-go/model"
    zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
    reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func main() {
	tracer, err := newTracer()
    if err != nil {
        log.Fatal(err)
    }

	router := mux.NewRouter()

	router.Handle("/metrics", promhttp.Handler())

	router.HandleFunc("/insurance/simpleQuote",  controllers.SimpleQuote).Methods("POST")

	router.Use(zipkinhttp.NewServerMiddleware(
        tracer,
        zipkinhttp.SpanName("request")), // name for request span
    )

	port := os.Getenv("PORT")
	if port == "" {
	    port = "8000" //localhost
	}

	log.Println(port)

	err = http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		log.Fatal(err)
	}
}

func newTracer() (*zipkin.Tracer, error) {
   // The reporter sends traces to zipkin server
   endpointURL := os.Getenv("ZIPKIN_ENDPOINT")
   reporter := reporterhttp.NewReporter(endpointURL)

   // Local endpoint represent the local service information
   localEndpoint := &model.Endpoint{ServiceName: "insurance_service", Port: 8000}

   // Sampler tells you which traces are going to be sampled or not. In this case we will record 100% (1.00) of traces.
   sampler, err := zipkin.NewCountingSampler(1)
   if err != nil {
      return nil, err
   }

   t, err := zipkin.NewTracer(
      reporter,
      zipkin.WithSampler(sampler),
      zipkin.WithLocalEndpoint(localEndpoint),
   )
   if err != nil {
      return nil, err
   }

   return t, err
}
