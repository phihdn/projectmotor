package db

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sqlx.DB

const DATABASE_URL string = "postgres://phihdn:S3cret@localhost:5432/projectmotor?search_path=public&sslmode=disable"

func OpenDB() error {
	conn, err := sqlx.Connect("pgx", DATABASE_URL)
	if err != nil {
		return err
	}
	DB = conn
	return nil
}

func CloseDB() error {
	return DB.Close()
}

// NOTE ->> only for testing, remove after actual interactions with database
func SetupDB() error {
	_, err := DB.Exec(`create table if not exists messages ("id" serial PRIMARY KEY, "message" text NOT NULL);`)
	if err != nil {
		return err
	}
	_, err = DB.Exec(`insert into messages (message) values ('Hello, world!');`)
	if err != nil {
		return err
	}
	return nil
}

func GetMessage() (string, error) {
	var message string
	err := DB.QueryRow("select message from messages limit 1;").Scan(&message)
	if err != nil {
		return "", err
	}
	return message, nil
}

// <<- NOTE
