package mysql

import (
	"database/sql"
	"fmt"
	"log"
)

type personEdit struct {
	ID        int     `db:"id"`
	Email     string  `db:"email"`
	FirstName string  `db:"firstName"`
	LastName  string  `db:"lastName"`
	Avatar    *string `db:"avatar"`
}

// albumsByArtist queries for albums that have the specified artist name.
func (service *mySqlService) GetPerson(id string) (personEdit, error) {
	// An albums slice to hold data from returned rows.
	var person personEdit

	row := service.db.QueryRow("SELECT * FROM person WHERE id = ?", id)
	if err := row.Scan(&person.ID, &person.Avatar, &person.Email, &person.FirstName, &person.LastName); err != nil {
		if err == sql.ErrNoRows {
			return person, fmt.Errorf("GetPerson %q: no such person", id)
		}
		return person, fmt.Errorf("GetPerson %q: %v", id, err)
	}
	return person, nil
}

// albumsByArtist queries for albums that have the specified artist name.
func (service *mySqlService) GetPeople(pageNum int) ([]personEdit, int, int, int, int, error) {
	// An albums slice to hold data from returned rows.
	var people []personEdit

	var totalCount int
	// Count total items
	countQuery := "SELECT COUNT(*) FROM person"
	err := service.db.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		log.Fatalf("Could not execute count query: %v", err)
	}

	// Calculate total pages
	pageSize := 10
	totalPages := totalCount / pageSize
	if totalCount%pageSize != 0 {
		totalPages++
	}

	offset := (pageNum - 1) * pageSize
	rows, err := service.db.Query("SELECT * FROM person ORDER BY id LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, -1, -1, -1, -1, fmt.Errorf("GetPeople : %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var person personEdit
		if err := rows.Scan(&person.ID, &person.Email, &person.FirstName, &person.LastName, &person.Avatar); err != nil {
			return nil, -1, -1, -1, -1, fmt.Errorf("GetPeople : %v", err)
		}
		people = append(people, person)
	}
	if err := rows.Err(); err != nil {
		return nil, -1, -1, -1, -1, fmt.Errorf("GetPeople : %v", err)
	}
	return people, pageNum, pageSize, totalPages, totalCount, nil
}
