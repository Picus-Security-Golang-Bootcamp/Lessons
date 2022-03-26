package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ready", ready)
	http.HandleFunc("/param", withQueryParam)
	http.HandleFunc("/form", withFormValue)
	http.HandleFunc("/create", userCreate)

	log.Println(http.ListenAndServe(":8080", nil))
}

//curl -v "localhost:8080/ready"
func ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("READY"))
}

//curl -v "localhost:8080/param?name=kaan"
func withQueryParam(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("name")

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(param))
}

//curl -v "localhost:8080/param?name=kaan"
func withFormValue(w http.ResponseWriter, r *http.Request) {
	param1 := r.FormValue("param1")
	param2 := r.FormValue("param2")

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(param1 + param2))
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	var u User
	w.Header().Set("Content-Type", "application/json")
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&u)
	//dec.DisallowUnknownFields()

	if r.Method != http.MethodPost {
		http.Error(w, "You can use POST method only", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, _ := json.Marshal(u)

	w.Write(resp)
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
