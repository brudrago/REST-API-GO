package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // 👈 registra o driver "postgres"
)

const (
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
	host     = "go_db" // se o Go rodar DENTRO do Docker, troca para "go_db"
	port     = 5432    // 👈 usa int aqui
)

func ConnectDB() (*sql.DB, error) {
	// OBS: tem espaço entre os campos e o port é %d (int)
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	log.Println("Successfully connected!", dbname)
	return db, nil
}
