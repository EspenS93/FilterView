package mysql

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type MySqlService interface {
	GetPerson(id string) (personEdit, error)
	GetPeople(page int) ([]personEdit, int, int, int, int, error)
}
type mySqlService struct {
	db *sql.DB
}

func New() MySqlService {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "godb",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	return &mySqlService{
		db: db,
	}
}
