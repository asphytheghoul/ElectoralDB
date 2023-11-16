//get and post requests for login
//once logged in, if aadhar id is part of table display the data from voter table for that aadhar id
//else redirect them to registration form. once they register, insert into voter table and login table
//change the sql file
//for candidate,party,voter,ECI official only limit visibility to those respectively.
//for ECI officials - make 4 forms - ECI form, Constituency form, Polling booth form, Election form
//add party name section on registration form
//if user logged in, delete that user's record from the respective table. redirect them to register screen
//whenever information button is clicked, call procedure to display the information from the respective table in a tabular format
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"example.com/voter"
	"log"
	"net/http"
	"github.com/rs/cors"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type FormData struct {
	AadharId         string `json:"aadharId"`
	FirstName        string `json:"firstName"`
	MiddleName       string `json:"middleName"`
	LastName         string `json:"lastName"`
	Gender           string `json:"gender"`
	Dob              string `json:"dob"`
	Age              string `json:"age"`
	State            string `json:"state"`
	PhoneNumber      string `json:"phoneNumber"`
	ConstituencyName string `json:"constituencyName"`
	PollingBoothId   string `json:"pollingBoothId"`
	VoterId          string `json:"voterId"`
	CandidateId      string `json:"candidateId"`
	PartyName        string `json:"partyName"`
	PartySymbol      string `json:"partySymbol"`
	PartyPresident   string `json:"president"`
	PartyFunds       string `json:"partyFunds"`
	HeadQuarters     string `json:"headquarters"`
	PartyLeader      string `json:"partyLeader"`
	SeatsWon         string `json:"seatsWon"`
	UserName         string `json:"userName"`
	Password         string `json:"password"`
	Role             string `json:"role"`
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func insertVoter(data FormData, hashedPassword string) error {
	// Insert data into the login table
	_, err := db.Exec(`
		INSERT INTO login (username, password, role)
		VALUES (?, ?, ?)`,
		data.UserName, hashedPassword, data.Role)
	if err != nil {
		return fmt.Errorf("failed to insert into login table: %v", err)
	}

	// retrieve all data from login table
	rows, err := db.Query("SELECT * FROM login")
	if err != nil {
		return fmt.Errorf("error executing SELECT query: %v", err)
	}
	defer rows.Close()

	// Iterate over the rows, sending the results to
	// standard out.
	for rows.Next() {
		var username string
		var password string
		var role string
		if err := rows.Scan(&username, &password, &role); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		fmt.Println(username, password, role)
	}

	return nil
}

func insertCandidate(data FormData, hashedPassword string) error {
	// Insert data into the login table
	_, err := db.Exec(`
		INSERT INTO login (username, password, role)
		VALUES (?, ?, ?)`,
		data.UserName, hashedPassword, data.Role)
	if err != nil {
		return fmt.Errorf("failed to insert into login table: %v", err)
	}

	// Insert data into the candidate table
	_, err = db.Exec(`
	INSERT INTO candidate (aadhar_id, first_name, last_name, middle_name, gender, dob, age,phone_no, constituency_name, voter_id)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	data.AadharId, data.FirstName, data.LastName, data.MiddleName, data.Gender, data.Dob, data.Age, data.PhoneNumber, data.ConstituencyName, data.VoterId)
if err != nil {
	return fmt.Errorf("failed to insert into candidate table: %v", err)
}

return nil
}

func insertParty(data FormData, hashedPassword string) error {
	// Insert data into the login table
	_, err := db.Exec(`
		INSERT INTO login (username, password, role)
		VALUES (?, ?, ?)`,
		data.UserName, hashedPassword, data.Role)
	if err != nil {
		return fmt.Errorf("failed to insert into login table: %v", err)
	}

	// Insert data into the party table
	_, err = db.Exec(`
	INSERT INTO party (party_name, party_symbol, party_president, party_funds, headquarters, party_leader, seats_won)
	VALUES (?, ?, ?, ?, ?, ?, ?)`,
	data.PartyName, data.PartySymbol, data.PartyPresident, data.PartyFunds, data.HeadQuarters, data.PartyLeader, data.SeatsWon)
if err != nil {
	return fmt.Errorf("failed to insert into party table: %v", err)
}

return nil
}


func handleRegistration(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data FormData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, passwordErr := hashPassword(data.Password)
	if passwordErr != nil {
		http.Error(rw, passwordErr.Error(), http.StatusInternalServerError)
		return
	}

	// Add logic to determine the role and insert into the appropriate table
	switch data.Role {
	case "voter":
		err = insertVoter(data, hashedPassword)
	case "candidate":
		// Add logic to insert into the candidate table
		err = insertCandidate(data, hashedPassword)

	case "party":
		// Add logic to insert into the party table
		err = insertParty(data, hashedPassword)
	default:
		err = fmt.Errorf("Invalid role: %s", data.Role)
	}

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(data)
}

func showTables() error {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "akash561@2910",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "electdb",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return fmt.Errorf("failed to ping database: %v", pingErr)
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

func main() {
	// Initialize the database connection
	cfg := mysql.Config{
		User:   "root",
		Passwd: "akash561@2910",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "electdb",
	}

	c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, // Allow your client's domain
        AllowCredentials: true,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
    })
	handler := c.Handler(http.DefaultServeMux)

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// Show tables in the database
	err = showTables()
	if err != nil {
		log.Fatal(err)
	}

	// Start the HTTP server
	http.HandleFunc("/register", handleRegistration)
	log.Println("Server is available at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
