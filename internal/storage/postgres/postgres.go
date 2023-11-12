package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/cactusfriday/go-url-shortener/internal/common"
	"github.com/cactusfriday/go-url-shortener/internal/storage/base"
	_ "github.com/lib/pq"
)

type SQLStorage struct {
	StorageURL *sql.DB
}

// Return URL or Error if not found
func (s *SQLStorage) GetUrl(shortUrl string) (string, error) {
	var url string
	selectStatement := `SELECT url FROM urls WHERE short_url = $1`

	if err := s.StorageURL.QueryRow(selectStatement, shortUrl).Scan(&url); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New(fmt.Sprint("No such record: ", shortUrl))
		}
		return "", errors.New(fmt.Sprint("Failed to get DB records: ", err))
	}

	return url, nil
}

func (s *SQLStorage) PutUrl(shortUrl string, url string) error {
	db_url, err := s.GetUrl(shortUrl)

	// If record already exists -> return nil
	if db_url != "" {
		return nil
	}

	insertStatement := `INSERT INTO urls (short_url, url) VALUES ($1, $2)`
	_, err = s.StorageURL.Exec(insertStatement, shortUrl, url)
	if err != nil {
		return errors.New(fmt.Sprint("Adding data error: ", err))
	}

	return nil
}

// SQL Storage constructor
func NewSqlStorage() base.IStorage {
	creds := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		common.CFG.DBHost,
		common.CFG.DBPort,
		common.CFG.DBUserName,
		common.CFG.DBPassword,
		common.CFG.DBName,
	)
	db, _ := sql.Open("postgres", creds)

	return &SQLStorage{
		StorageURL: db,
	}
}
