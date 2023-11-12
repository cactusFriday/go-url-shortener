package storage

import (
	"github.com/cactusfriday/go-url-shortener/internal/storage/base"
	"github.com/cactusfriday/go-url-shortener/internal/storage/postgres"
	"github.com/cactusfriday/go-url-shortener/internal/storage/ram"
)

var AppStorage base.IStorage

func NewStorage(storageType string) base.IStorage {
	if AppStorage != nil {
		return AppStorage
	}
	switch storageType {
	case "sql":
		AppStorage = postgres.NewSqlStorage()
		return AppStorage
	case "ram":
		AppStorage = ram.NewRamStorage()
		return AppStorage
	}
	return nil
}
