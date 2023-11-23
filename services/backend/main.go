package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	PhoneNumber      string `json:"phone"`
	ConstituencyName string `json:"constituency"`
	PollingBoothId   string `json:"pollingBoothId"`
	VoterId          string `json:"voterId"`
	CandidateId      string `json:"candidateId"`
	PartyRep         string `json:"partyRep"`
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
	ElectionId		 string `json:"electionId"`
	ElectionType	 string `json:"electionType"`
	DateOfElection   string `json:"electionDate"`
	Seats		     string `json:"seats"`
	OfficialId       string `json:"officialId"`
	ConstituencyAssigned string `json:"constituencyAssigned"`
	PollBoothAssigned string `json:"pollBoothAssigned"`
	HigherRankId string `json:"higherRankId"`
	OfficialRank string `json:"officialRank"`
	PartyMemberCount string `json:"partyMemberCount"`
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func handleLogin(rw http.ResponseWriter,r *http.Request){
	// parse aadhar id and password from the request body
	var data FormData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	//check voter, candidate, party, eci official tables for aadhar id
	//if found, verify password and return user's role and details
	//else return error message
	var role string
	var aadhar_id string
	var password string
	var party_name string
	rows, err := db.Query("SELECT * FROM login WHERE aadhar_id = ?", data.AadharId)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&aadhar_id, &password, &role, &party_name); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

        err = bcrypt.CompareHashAndPassword([]byte(password), []byte(data.Password))
        if err != nil {
			log.Println(data.Password)
			log.Println(err)
            http.Error(rw, "Invalid password", http.StatusBadRequest)
            return
        }

	}
	//if aadhar id not found, return error message
	if aadhar_id == "" {
		http.Error(rw, "Aadhar id not found", http.StatusBadRequest)
		return
	}

	type ResponseData struct {
		AadharId   string `json:"aadhar_id"`
		Role       string `json:"role"`
		PartyName  string `json:"party_name"`
	}
	
	
	responseData := ResponseData{
		AadharId:  aadhar_id,
		Role:      role,
		PartyName: party_name,
	}
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(responseData)
	fmt.Print(responseData)
}

func insertUser(data FormData, hashedPassword string) error {
	_, err := db.Exec(`
		INSERT INTO login (aadhar_id, password, role, party_name)
		VALUES (?, ?, ?, ?)`,
		data.AadharId, hashedPassword, data.Role, data.PartyName)
	if err != nil {
		return fmt.Errorf("failed to insert into login table: %v", err)
	}

	rows, err := db.Query("SELECT * FROM login")
	if err != nil {
		return fmt.Errorf("error executing SELECT query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var aadharID string
		var password string
		var role string
		var partyName string
		if err := rows.Scan(&aadharID, &password, &role, &partyName); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		fmt.Println(aadharID, password, role, partyName)
	}

	return nil
}

func insertVoter(data FormData) error {
    // Insert data into the voter table
    _, err := db.Exec(`
        INSERT INTO voter (aadhar_id, first_name, last_name, middle_name, gender, dob, phone_no, state, constituency_name, polling_booth_id)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        data.AadharId, data.FirstName, data.LastName, data.MiddleName, data.Gender, data.Dob, data.PhoneNumber, data.State, data.ConstituencyName, data.PollingBoothId)
    if err != nil {
        return fmt.Errorf("failed to insert into voter table: %v", err)
    }
    return nil
}

func insertCandidate(data FormData) error {
    // Insert data into the candidate table
    _, err := db.Exec(`
        INSERT INTO candidate (aadhar_id, first_name, last_name, middle_name, gender, dob, phone_no, cons_fight,party_rep)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        data.AadharId, data.FirstName, data.LastName, data.MiddleName, data.Gender, data.Dob, data.PhoneNumber, data.ConstituencyName,data.PartyRep)
    if err != nil {
        return fmt.Errorf("failed to insert into candidate table: %v", err)
    }
    return nil
}

func insertParty(data FormData) error {
    // Insert data into the party table
    _, err := db.Exec(`
        INSERT INTO party (party_name, party_symbol, party_president, party_funds, headquarters, party_leader, seats_won)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        data.PartyName, data.PartySymbol, data.PartyPresident, data.PartyFunds, data.HeadQuarters, data.PartyLeader, data.SeatsWon)
    if err != nil {
        return fmt.Errorf("failed to insert into party table: %v", err)
    }
    return nil
}
func getConstDeets(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("call getconsdets()")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer rows.Close()

    var results []map[string]string

    for rows.Next() {
        var constituencyName string
        var maleCount string
        var femaleCount string
        var pollBoothCount string

        if err := rows.Scan(&constituencyName, &maleCount, &femaleCount, &pollBoothCount); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        row := map[string]string{
            "constituencyName": constituencyName,
            "maleCount":        maleCount,
            "femaleCount":      femaleCount,
            "pollBoothCount":   pollBoothCount,
        }

        results = append(results, row)
    }

    jsonData, err := json.Marshal(results)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    errorObject := map[string]string{"error": message}
    jsonData, _ := json.Marshal(errorObject)
    w.Write(jsonData)
}

func getvoterinformation(w http.ResponseWriter, r *http.Request) {
    aadharId := r.URL.Query().Get("aadharId")
    if aadharId == "" {
        http.Error(w, "Aadhar ID is required", http.StatusBadRequest)
        return
    }
    row := db.QueryRow("SELECT * FROM voter WHERE aadhar_id = ?", aadharId)

    var aadhar_id string
    var first_name string
    var last_name string
    var middle_name sql.NullString
    var gender string
    var dob string
    var age string
    var state string
    var phone_no string
    var constituency_name string
    var poll_booth_id string
    var voter_id string

    if err := row.Scan(&aadhar_id, &first_name, &last_name, &middle_name, &gender, &dob, &age, &state, &phone_no, &constituency_name, &poll_booth_id, &voter_id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    voter := map[string]string{
        "aadharId": aadhar_id,
        "firstName": first_name,
        "lastName": last_name,
        "middleName": middle_name.String,
        "gender":gender,
        "dob":dob,
        "age":age,
        "state":state,
        "phoneNumber":phone_no,
        "constituencyName":constituency_name,
        "pollingBoothId":poll_booth_id,
        "voterId":voter_id,
    }

    jsonData, err := json.Marshal(voter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

func getcandidateinformation(w http.ResponseWriter, r *http.Request) {
    aadharId := r.URL.Query().Get("aadharId")
    if aadharId == "" {
        http.Error(w, "Aadhar ID is required", http.StatusBadRequest)
        return
    }
    row := db.QueryRow("SELECT * FROM candidate WHERE aadhar_id = ?", aadharId)

    var aadharID string
    var firstName string
    var lastName string
    var middleName sql.NullString
    var gender string
    var dob string
    var age string
    var phoneNumber string
    var constituencyFighting string
    var candidateID string
    var partyRep string

    if err := row.Scan(&aadharID, &firstName, &lastName, &middleName, &gender, &dob, &age, &phoneNumber, &constituencyFighting, &candidateID, &partyRep); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    candidate := map[string]string{
        "aadharId": aadharID,
        "firstName": firstName,
        "lastName": lastName,
        "middleName": middleName.String,
        "gender": gender,
        "dob": dob,
        "age": age,
        "phoneNumber": phoneNumber,
        "constituencyFighting": constituencyFighting,
        "candidateId": candidateID,
        "partyRep": partyRep,
    }

    jsonData, err := json.Marshal(candidate)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

func getpartyinformation(w http.ResponseWriter, r *http.Request) {
    partyName := r.URL.Query().Get("partyName")
    if partyName == "" {
        http.Error(w, "Party name is required", http.StatusBadRequest)
        return
    }
    row := db.QueryRow("SELECT party_name, party_symbol, president, party_funds, headquarters, seats_won, party_member_count FROM party where party_name = ?", partyName)
    var partySymbol string
    var president string
    var partyFunds string
    var headquarters string
    var seatsWon sql.NullString
    var partyMemberCount string

    if err := row.Scan(&partyName, &partySymbol, &president, &partyFunds,&headquarters,&seatsWon,&partyMemberCount); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    party := map[string]string{
        "partyName": partyName,
        "partySymbol": partySymbol,
        "president": president,
        "partyFunds": partyFunds,
        "headquarters": headquarters,
        "seatsWon": seatsWon.String,
        "partyMemberCount": partyMemberCount,
    }

    jsonData, err := json.Marshal(party)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

func getofficialinformation(w http.ResponseWriter, r *http.Request) {
    aadharId := r.URL.Query().Get("aadharId")
    fmt.Println(aadharId)
    if aadharId == "" {
        http.Error(w, "Aadhar ID is required", http.StatusBadRequest)
        return
    }
    row := db.QueryRow("SELECT aadhar_id, first_name, last_name, middle_name, gender, dob, age, phone_no, constituency_assigned, poll_booth_assigned, official_id, official_rank, higher_rank_id FROM official WHERE aadhar_id = ?", aadharId)

    var aadharID string
    var firstName string
    var lastName string
    var middleName string
    var gender string
    var dob string
    var age string
    var phoneNumber string
    var constituencyAssigned string
    var pollBoothAssigned sql.NullString
    var officialID string
    var officialRank string
    var higherRankID string

    if err := row.Scan(&aadharID, &firstName, &lastName, &middleName, &gender, &dob, &age, &phoneNumber, &constituencyAssigned, &pollBoothAssigned, &officialID, &officialRank, &higherRankID); err != nil {
        errorResponse := map[string]string{"error": err.Error()}
        jsonData, _ := json.Marshal(errorResponse)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(jsonData)
        fmt.Println("here")
        return
    }

    official := map[string]string{
        "aadharId": aadharID,
        "firstName": firstName,
        "lastName": lastName,
        "middleName": middleName,
        "gender": gender,
        "dob": dob,
        "age": age,
        "phoneNumber": phoneNumber,
        "constituencyAssigned": constituencyAssigned,
        "pollBoothAssigned": pollBoothAssigned.String,
        "officialId": officialID,
        "officialRank": officialRank,
        "higherRankId": higherRankID,
    }

    jsonData, err := json.Marshal(official)
    if err != nil {
        errorResponse := map[string]string{"error": err.Error()}
        jsonData, _ := json.Marshal(errorResponse)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(jsonData)
        fmt.Println("here 2")
        return
    }
    fmt.Println(jsonData)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
    fmt.Println("here 3")
}

func UpdateelectionDetails(db *sql.DB, eleiD string, eletyp string, dateofe string, seats string) error {
	fmt.Println("eleiD:", eleiD, "eletyp:", eletyp, "dateofe:", dateofe, "seats:", seats)
	query := "UPDATE election SET election_type = ?, seats = ?, dateofelection = ?  WHERE election_id = ?"
	_, err := db.Exec(query, eletyp, seats, dateofe, eleiD)
	if err != nil {
		return err
	}
	return nil
}

func DeleteelectionByID(db *sql.DB, eleid string) error {
	query := "DELETE FROM election WHERE election_id = ?"
	_, err := db.Exec(query, eleid)
	if err != nil {
		return err
	}
	return nil
}

func handleUpdateElection(w http.ResponseWriter, r *http.Request) {
	// Parse FormData from request body
	var data FormData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call UpdateelectionDetails function
	err = UpdateelectionDetails(db, data.ElectionId, data.ElectionType, data.DateOfElection, data.Seats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Election updated successfully"))
}

func handleDeleteElection(w http.ResponseWriter,r *http.Request){
		// Parse FormData from request body
		var data FormData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		// Call UpdateelectionDetails function
		err = DeleteelectionByID(db, data.ElectionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		// Send success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Election Deleted successfully"))
}

// Update function for Voter
func UpdateVoterDetails(db *sql.DB, aadharID string, firstName string, lastName string, middleName string, gender string, dob string, state string, phoneNo string, voterID string) error {
    query := "UPDATE voter SET voter_id = ?, first_name = ?, last_name = ?, middle_name = ?, gender = ?, dob = ?, state = ?, phone_no = ? WHERE aadhar_id = ?"
    _, err := db.Exec(query, voterID, firstName, lastName, middleName, gender, dob, state, phoneNo, aadharID)
    if err != nil {
        return err
    }
    return nil
}

// Update function for Party
func UpdatePartyDetails(db *sql.DB, partyName string, partySymbol string, president string, partyFunds string, headquarters string, partyMemberCount string) error {
    query := "UPDATE party SET party_symbol = ?, president = ?, party_funds = ?, headquarters = ? party_member_count = ?WHERE party_name = ?"
    _, err := db.Exec(query, partySymbol, president, partyFunds, headquarters,partyMemberCount, partyName)
    if err != nil {
        return err
    }
    return nil
}

// Update function for Candidate
func UpdateCandidateDetails(db *sql.DB, aadharID string, firstName string, lastName string, middleName string, gender string, dob string, phoneNo string, constituencyFighting string, candidateId string, partyRep string) error {
    query := "UPDATE candidate SET first_name = ?, last_name = ?, middle_name = ?, gender = ?, dob = ?, phone_no = ?, cons_fight = ?, candidate_id = ?, party_rep = ? WHERE aadhar_id = ?"
    _, err := db.Exec(query, firstName, lastName, middleName, gender, dob, phoneNo,constituencyFighting,candidateId,partyRep, aadharID)
    if err != nil {
        return err
    }
    return nil
}

// Update function for Official
func UpdateOfficialDetails(db *sql.DB, aadharID string, firstName string, lastName string, middleName string, gender string, dob string, phoneNo string, constituencyAssigned string, pollBoothAssigned string, officialID string) error {
    query := "UPDATE official SET aadhar_id = ?, first_name = ?, last_name = ?, middle_name = ?, gender = ?, dob = ?, phone_no = ?, constituency_assigned = ?, poll_booth_assigned = ? WHERE official_id = ?"
    _, err := db.Exec(query, aadharID, firstName, lastName, middleName, gender, dob, phoneNo, constituencyAssigned, pollBoothAssigned, officialID)
    if err != nil {
        return err
    }
    return nil
}

func DeleteVoterByAadharID(db *sql.DB, aadharID string) error {
    query := "DELETE FROM voter WHERE aadhar_id = ?"
    _, err := db.Exec(query, aadharID)
    if err != nil {
        return err
    }
    return nil
}

func DeleteCandidateByAadharID(db *sql.DB, aadharID string) error {
	query := "DELETE FROM candidate WHERE aadhar_id = ?"
	_, err := db.Exec(query, aadharID)
	if err != nil {
		return err
	}
	return nil
}

func DeletePartyByPartyName(db *sql.DB, partyName string) error {
	query := "DELETE FROM party WHERE party_name = ?"
	_, err := db.Exec(query, partyName)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOfficialByAadharID(db *sql.DB, aadharID string) error {
	query := "DELETE FROM official WHERE aadhar_id = ?"
	_, err := db.Exec(query, aadharID)
	if err != nil {
		return err
	}
	return nil
}

func handleDeleteVoter(w http.ResponseWriter, r *http.Request) {
	// Parse FormData from request body
	var data FormData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call DeleteVoterByAadharID function
	err = DeleteVoterByAadharID(db, data.AadharId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Voter deleted successfully"))
}

func handleDeleteCandidate(w http.ResponseWriter, r *http.Request) {
	// Parse FormData from request body
	var data FormData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call DeleteCandidateByAadharID function
	err = DeleteCandidateByAadharID(db, data.AadharId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Candidate deleted successfully"))
}

func handleDeleteParty(w http.ResponseWriter, r *http.Request) {
	// Parse FormData from request body
	var data FormData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call DeletePartyByPartyName function
	err = DeletePartyByPartyName(db, data.PartyName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Party deleted successfully"))
}

func handleDeleteOfficial(w http.ResponseWriter, r *http.Request) {
	// Parse FormData from request body
	var data FormData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call DeleteOfficialByAadharID function
	err = DeleteOfficialByAadharID(db, data.AadharId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Official deleted successfully"))
}

func handleUpdateVoter(w http.ResponseWriter, r *http.Request) {
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = UpdateVoterDetails(db, data.AadharId, data.FirstName, data.LastName, data.MiddleName, data.Gender, data.Dob, data.State, data.PhoneNumber, data.VoterId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Voter updated successfully"))
}

func handleUpdateParty(w http.ResponseWriter, r *http.Request) {
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = UpdatePartyDetails(db, data.PartyName, data.PartySymbol, data.PartyPresident, data.PartyFunds, data.HeadQuarters, data.PartyMemberCount)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Party updated successfully"))
}

func handleUpdateCandidate(w http.ResponseWriter, r *http.Request) {
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = UpdateCandidateDetails(db, data.AadharId, data.FirstName, data.LastName, data.MiddleName, data.Gender, data.Dob, data.PhoneNumber, data.ConstituencyName, data.CandidateId, data.PartyRep)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Candidate updated successfully"))
}

func handleUpdateOfficial(w http.ResponseWriter, r *http.Request) {
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = UpdateOfficialDetails(db, data.AadharId, data.FirstName, data.LastName, data.MiddleName, data.Gender, data.Dob, data.PhoneNumber, data.ConstituencyAssigned, data.PollBoothAssigned, data.OfficialId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Official updated successfully"))
}

func selectVoter(data FormData) (map[string]string, error) {
    rows,err := db.Query("SELECT * FROM voter where aadhar_id = ?", data.AadharId)
    if err != nil {
        return nil, fmt.Errorf("error executing SELECT query: %v", err)
    }
    defer rows.Close()

    var voter map[string]string

    for rows.Next() {
        var aadharID string
        var firstName string
        var lastName string
        var middleName string
        var gender string
        var dob string
        var age string
        var state string
        var phoneNumber string
        var constituencyName string
        var pollBoothID string
        var voterID string

        if err := rows.Scan(&aadharID, &firstName, &lastName, &middleName, &gender, &dob, &age, &state,&phoneNumber, &constituencyName, &pollBoothID, &voterID); err != nil {
            return nil, fmt.Errorf("error scanning row: %v", err)
        }

        voter = map[string]string{
            "aadharId": aadharID,
            "firstName": firstName,
            "lastName": lastName,
            "middleName": middleName,
            "gender": gender,
            "dob": dob,
            "age": age,
            "state": state,
            "phoneNumber": phoneNumber,
            "constituencyName": constituencyName,
            "pollingBoothId": pollBoothID,
            "voterId": voterID,
        }
    }

    return voter, nil
}

func selectOfficials(data FormData) error {
	rows, err := db.Query("SELECT * FROM official where aadhar_id = ?", data.AadharId)
	if err != nil {
		return fmt.Errorf("error executing SELECT query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var aadharID string
		var firstName string
		var lastName string
		var middleName string
		var gender string
		var dob string
		var age string
		var phoneNumber string
		var constituencyAssigned string
		var pollBoothAssigned string
		var officialID string 
		var officialRank string
		var higherRankID string
		if err := rows.Scan(&aadharID, &firstName, &lastName, &middleName, &gender, &dob, &age,&phoneNumber, &constituencyAssigned, &pollBoothAssigned, &officialID,&officialRank,&higherRankID); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}

	}

	return nil
}

func selectCandidate(data FormData) error {
	rows,err := db.Query("SELECT * FROM candidate where aadhar_id = ?", data.AadharId)
	if err != nil {
		return fmt.Errorf("error executing SELECT query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var aadharID string
		var firstName string
		var lastName string
		var middleName string
		var gender string
		var dob string
		var age string
		var phoneNumber string
		var constituencyFighting string
		var candidateID string
		var partyRep string

		if err := rows.Scan(&aadharID, &firstName, &lastName, &middleName, &gender, &dob, &age, &phoneNumber, &constituencyFighting, &candidateID, &partyRep); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
	}
	return nil
}

func selectParty(data FormData) error {
	rows,err := db.Query("SELECT * FROM party where party_name = ?", data.PartyName)
	if err != nil {
		return fmt.Errorf("error executing SELECT query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var partyName string
		var partySymbol string
		var partyPresident string
		var partyFunds string
		var headquarters string
		var partyLeader string
		var seatsWon string

		if err := rows.Scan(&partyName, &partySymbol, &partyPresident, &partyFunds, &headquarters, &partyLeader, &seatsWon); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
	}
	return nil
}


func handleInsertVoter(w http.ResponseWriter, r *http.Request) {
    // Parse FormData from request body
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call insertVoter function
    err = insertVoter(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("voter inserted successfully"))
}

func handleInsertCandidate(w http.ResponseWriter, r *http.Request) {
    // Parse FormData from request body
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Call insertCandidate function
    err = insertCandidate(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Candidate inserted successfully"))
}

func handleInsertParty(w http.ResponseWriter, r *http.Request) {
    // Parse FormData from request body
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call insertParty function
    err = insertParty(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Party inserted successfully"))
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

	hashedPassword, passwordErr := hashPassword(data.Password)
	if passwordErr != nil {
		http.Error(rw, passwordErr.Error(), http.StatusInternalServerError)
		return
	}
	switch data.Role {
	case "voter":
		err = insertUser(data, hashedPassword)
	case "candidate":
		err = insertUser(data, hashedPassword)
	case "party":
		err = insertUser(data, hashedPassword)
    case "official":
        err = insertUser(data, hashedPassword)
	default:
		err = fmt.Errorf("invalid role: %s", data.Role)
	}

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(data)
}

func handleSelectVoter(w http.ResponseWriter, r *http.Request) {
    // Parse FormData from request body
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call selectVoter function
    voter, err := selectVoter(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send success response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(voter)
}

func handleSelectOfficial(w http.ResponseWriter, r *http.Request) {
    // Parse FormData from request body
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call selectVoter function
    err = selectOfficials(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send success response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Official selected successfully"})
}


func handleSelectCandidate(w http.ResponseWriter, r *http.Request) {
    // Parse FormData from request body
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call selectCandidate function
    err = selectCandidate(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send success response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Voter selected successfully"})
}

func handleSelectParty(w http.ResponseWriter, r *http.Request) {
    var data FormData
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = selectParty(data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send success response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Voter selected successfully"})
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

	err = showTables()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/register", handleRegistration)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/insertVoter", handleInsertVoter)
	http.HandleFunc("/insertCandidate", handleInsertCandidate)
	http.HandleFunc("/insertParty", handleInsertParty)
	http.HandleFunc("/selectVoter", handleSelectVoter)
	http.HandleFunc("/selectCandidate", handleSelectCandidate)
	http.HandleFunc("/selectParty", handleSelectParty)
	http.HandleFunc("/selectOfficials",handleSelectOfficial)
	http.HandleFunc("/getConstDeets", getConstDeets)
	http.HandleFunc("/getvoterinformation", getvoterinformation)
	http.HandleFunc("/getcandidateinformation", getcandidateinformation)
	http.HandleFunc("/getpartyinformation", getpartyinformation)
	http.HandleFunc("/getofficialinformation", getofficialinformation)
	http.HandleFunc("/updateVoter", handleUpdateVoter)
	http.HandleFunc("/updateParty", handleUpdateParty)
	http.HandleFunc("/updateCandidate", handleUpdateCandidate)
	http.HandleFunc("/updateOfficial", handleUpdateOfficial)
	http.HandleFunc("/deleteVoter", handleDeleteVoter)
	http.HandleFunc("/deleteCandidate", handleDeleteCandidate)
	http.HandleFunc("/deleteParty", handleDeleteParty)
	http.HandleFunc("/deleteOfficial", handleDeleteOfficial)
	http.HandleFunc("/updateElection",handleUpdateElection)
	http.HandleFunc("/deleteElection",handleDeleteElection)
	log.Println("Server is available at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
