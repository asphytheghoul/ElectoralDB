package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
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
	CandidateId  string `json:"candidateId"`
	PartyName   string `json:"partyName"`
	PartySymbol string `json:"partySymbol"`
	PartyPresident    string `json:"president"`
	PartyFunds string `json:"partyFunds"`
	HeadQuarters string `json:"headquarters"`
	PartyLeader string `json:"partyLeader"`
	SeatsWon string `json:"seatsWon"`
}

func showTables() error {
	connStr := "username:password@tcp(127.0.0.1:3306)/electdb"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("unable to connect to the database: %v", err)
	}

	rows, err := db.Query("SHOW TABLES;")
	if err != nil {
		return fmt.Errorf("error executing SHOW TABLES query: %v", err)
	}
	defer rows.Close()

	fmt.Println("Tables in the database:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return fmt.Errorf("error scanning table name: %v", err)
		}
		fmt.Println(tableName)
	}

	return nil
}

func insertFormDataIntoDatabase(data FormData) error {
	connStr := "root:akash561@2910@tcp(127.0.0.1:3306)/electdb"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("unable to connect to the database: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback() 

	switch {
	case data.VoterId != "":
		_, err = tx.Exec(`
			INSERT INTO voter (aadhar_id, first_name, last_name, middle_name, gender, dob, age, state, phone_no, constituency_name, poll_booth_id, voter_id)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			data.AadharId, data.FirstName, data.LastName, data.MiddleName, data.Gender, data.Dob, data.Age, data.State, data.PhoneNumber, data.ConstituencyName, data.PollingBoothId, data.VoterId)
	case data.CandidateId != "":
		_, err = tx.Exec(`
			INSERT INTO candidate (aadhar_id, first_name, last_name, middle_name, gender, dob, age, phone_no, cons_fight, candidate_id, party_rep)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			data.AadharId, data.FirstName, data.LastName, data.MiddleName, data.Gender, data.Dob, data.Age, data.PhoneNumber, data.ConstituencyName, data.CandidateId, data.PartyName)
		_, err = tx.Exec(`UPDATE party SET party_member_count = party_member_count + 1 WHERE party_name = ?`, data.PartyName)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
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
			fmt.Println(fmt.Sprintf("%+v", data))

			err = insertFormDataIntoDatabase(data)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}

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
	err := showTables()
	if err != nil {
		log.Fatal(err)
	}
}

// func main() {
// 	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		rw.Header().Set("Access-Control-Allow-Origin", "*")
// 		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 		if r.Method == http.MethodPost {
// 			var data FormData
// 			err := json.NewDecoder(r.Body).Decode(&data)
// 			if err != nil {
// 				http.Error(rw, err.Error(), http.StatusBadRequest)
// 				return
// 			}
// 			fmt.Println(fmt.Sprintf("%+v", data))

// 			rw.Header().Set("Content-Type", "application/json")
// 			rw.WriteHeader(http.StatusOK)
// 			json.NewEncoder(rw).Encode(data)
// 		} else {
// 			resp := []byte(`{"status":"ok"}`)
// 			rw.Header().Set("Content-Type", "application/json")
// 			rw.Header().Set("Content-Length", fmt.Sprint(len(resp)))
// 			rw.Write(resp)
// 		}
// 	})

// 	log.Println("Server is available at http://localhost:8000")
// 	log.Fatal(http.ListenAndServe(":8000", handler))
// }
