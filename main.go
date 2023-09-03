package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type employee struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Title string `json:"Title"`
}

type allEmployees []employee

var employees = allEmployees{
	{
		ID:    "1",
		Name:  "Joe Marcus",
		Title: "Product Manager",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome home!")
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee employee
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter data with the employee name and job title")
	}

	json.Unmarshal(reqBody, &newEmployee)
	employees = append(employees, newEmployee)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEmployee)
}

func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/employee", createEmployee).Methods("POST")
	router.HandleFunc("/employees", getAllEmployees).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
