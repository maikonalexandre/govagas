package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitDB() error {
	var err error
	db, err = InitSqLite()

	if err != nil {
		return fmt.Errorf("error initialize sqlite %v", err)
	}

	return nil
}

func GetSqLiteDb() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
