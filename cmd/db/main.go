package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func fmtDsn(host, port, user, dbName, password string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host,
		port,
		user,
		dbName,
		password,
	)
}

func dbExec(dsn string, f func(*sql.DB) error) {
	db, err := sql.Open(
		os.Getenv("POSTGRES_DRIVER"),
		dsn,
	)

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("sql.Close error %s", err)
		}
	}()

	if err != nil {
		log.Printf("sql.Open error %s", err)
	}

	if err = f(db); err != nil {
		log.Printf("sql.Query error %s", err)
	}
}

func createDatabase() {
	dsn := fmtDsn(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	f := func(db *sql.DB) error {
		query := "CREATE DATABASE panforyou"
		return db.QueryRow(query).Err()
	}
	dbExec(dsn, f)
}

func createTable() {
	dsn := fmtDsn(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	f := func(db *sql.DB) error {
		query := "CREATE TABLE breads (id text,name text,createdAt TIMESTAMP,PRIMARY KEY(id));"
		return db.QueryRow(query).Err()
	}
	dbExec(dsn, f)
}

func main() {
	createDatabase()
	createTable()
}
