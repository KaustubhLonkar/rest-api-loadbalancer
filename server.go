package main

import (
	"net/http"

	lb "github.appl.ge.com/geappliancesales/loadbalancer/loadbalancer"
)

func main() {
	handler := http.NewServeMux()
	///we create a new router to expose our api
	//to our users

	handler.HandleFunc("/api/location", lb.GetLocation)
	//Every time a  request is sent to the endpoint ("/api/location")
	//the function GetLocation will be invoked
	http.ListenAndServe("0.0.0.0:8080", handler)
	//we tell our api to listen to all request to port 8080.

}
