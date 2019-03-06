package main

import (
	"encoding/json"
	"log"
	"net/http"
	
	 "gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	
	. "repository"
	
	. "models"
	//. "github.com/rest-api/models"
)

var ss = SingularServer{}

// GET list of patients
func AllPatientEndPoint(w http.ResponseWriter, r *http.Request) {

	patiente, err := ss.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, patiente)
}

// POST a new movie
func CreatePatientEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var patiente Asu
	if err := json.NewDecoder(r.Body).Decode(&patiente); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	patiente.ID = bson.NewObjectId()
	if err := ss.Insert(patiente); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, patiente)//TODO : find our why respond with json is used here
}

func SinglePatientEndpoint(w http.ResponseWriter, r *http.Request) {
	// GET a patient by its ID
	params := mux.Vars(r)
	patiente, err := ss.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Patient ID")
		return
	}
	respondWithJson(w, http.StatusOK, patiente)

}


func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)//this response is our json data of structure Asu 
}

func init() {

	ss.Server = "localhost"
	ss.Database = "patientsdb"
	ss.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/patients", AllPatientEndPoint).Methods("GET")
	r.HandleFunc("/patients/{id}", CreatePatientEndPoint).Methods("POST")
	r.HandleFunc("/patient/{id}", SinglePatientEndpoint).Methods("GET")

	
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

