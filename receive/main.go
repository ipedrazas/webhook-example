package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hookHandler)
	log.Println("Listening in 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err.Error())
	}
	log.Println("Body received:")
	log.Println(string(body))
	var arbitraryJson map[string]interface{}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.Unmarshal(body, &arbitraryJson); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	log.Println("Object from json received")
	log.Println(arbitraryJson)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(arbitraryJson); err != nil {
		panic(err)
	}

}
