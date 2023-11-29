package config

import (
	"os"

	"github.com/maikonalexandre/govagas/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite:")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database not found, creating...")

		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errf("sqlite opening error %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errf("sqlite automigration error %v", err)
		return nil, err
	}

	return db, nil
}
