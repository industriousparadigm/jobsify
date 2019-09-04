package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type company struct {
	ID          int    `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Location    string `json:"Location"`
}

type allCompanies []company

var companies = allCompanies{
	{
		ID:          1,
		Name:        "Lego",
		Description: "Toy company building connectable bricks that form bigger shapes",
		Location:    "London",
	},
}

func main() {
	// initCompanies()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/company", createCompany).Methods("POST")
	router.HandleFunc("/companies", getAllCompanies).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home.")
}

func createCompany(w http.ResponseWriter, r *http.Request) {
	var newCompany company
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter company name, description and location")
	}
	json.Unmarshal(reqBody, &newCompany)
	companies = append(companies, newCompany)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCompany)
}

func getAllCompanies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(companies)
}
