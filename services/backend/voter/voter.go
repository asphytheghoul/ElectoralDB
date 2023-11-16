package voter

import (
	"database/sql"
)

func Selectvoterdetails(db *sql.DB, voterid string) ([]map[string]interface{}, error) {

	rows, err := db.Query("select voter_id,first_name from voter where voter_id = ?", voterid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []map[string]interface{}

	for rows.Next() {

		var (
			vid   string
			fname string
		)
		err := rows.Scan(&vid, &fname)
		if err != nil {
			return nil, err
		}

		rowData := map[string]interface{}{
			"Voter_ID":   vid,
			"Voter_name": fname}
		data = append(data, rowData)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func InsertVoterDetails(db *sql.DB, voterID string, firstName string) error {
	// Example INSERT query
	query := "INSERT INTO voter (voter_id, first_name) VALUES (?, ?)"

	// Execute the query to insert data
	_, err := db.Exec(query, voterID, firstName)
	if err != nil {
		return err
	}

	// No need to defer rows.Close() or handle returned rows in case of INSERT operation

	return nil
}

func UpdateVoterDetails(db *sql.DB, voterID string, newFirstName string) error {
	// Example UPDATE query
	query := "UPDATE voter SET first_name = ? WHERE voter_id = ?"

	// Execute the query to update data
	_, err := db.Exec(query, newFirstName, voterID)
	if err != nil {
		return err
	}

	// No need to defer rows.Close() or handle returned rows in case of UPDATE operation

	return nil
}

func DeleteVoterByID(db *sql.DB, voterID string) error {
	// Example DELETE query
	query := "DELETE FROM voter WHERE voter_id = ?"

	// Execute the query to delete data
	_, err := db.Exec(query, voterID)
	if err != nil {
		return err
	}

	// No need to defer rows.Close() or handle returned rows in case of DELETE operation

	return nil
}
