package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
)


// Must receive a body of number of coches and number of semaforos
// Must return a json array of created coches
func returncars(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(coches)
}


func handleInput(w http.ResponseWriter, r *http.Request){
	for k, v := range r.URL.Query(){
		fmt.Printf("%s: %s\n", k, v)
	}


	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.WriteHeader(http.StatusAccepted)
}

func initServer(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/load", returncars).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(":3000", router))
}