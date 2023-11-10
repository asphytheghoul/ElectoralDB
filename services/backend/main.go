package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FormData struct {
	AadharId     string `json:"aadharId"`
	FirstName    string `json:"firstName"`
	MiddleName   string `json:"middleName"`
	LastName     string `json:"lastName"`
	Gender       string `json:"gender"`
	Dob          string `json:"dob"`
	Age          string `json:"age"`
	State        string `json:"state"`
	PhoneNumber  string `json:"phoneNumber"`
	ConstituencyName string `json:"constituencyName"`
	PollingBoothId string `json:"pollingBoothId"`
	VoterId      string `json:"voterId"`
}

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == http.MethodPost {
			var data FormData
			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}
			// Now you can use the data
			fmt.Println(fmt.Sprintf("%+v", data))

			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(data)
		} else {
			resp := []byte(`{"status":"ok"}`)
			rw.Header().Set("Content-Type", "application/json")
			rw.Header().Set("Content-Length", fmt.Sprint(len(resp)))
			rw.Write(resp)
		}
	})

	log.Println("Server is available at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
