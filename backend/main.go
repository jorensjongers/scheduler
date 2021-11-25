/*
 * Scheduler API
 *
 * Basic API for retrieving schedules, based on a ScheduleInput object.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/DennisMaes6/scheduler/backend/server"
)

func main() {
	log.Printf("Server started")

	DefaultApiService := openapi.NewDefaultApiService()
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	log.Printf("Here")
	router := openapi.NewRouter(DefaultApiController)
	log.Printf("Here2")
	
	log.Fatal(http.ListenAndServe(":8080", router))
	log.Printf("Here3")
}
