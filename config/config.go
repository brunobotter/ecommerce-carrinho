package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	//Initialize SQLite
	db, err = InitializeSql()
	if err != nil {
		return fmt.Errorf("Error initialize mysql: %v", err)
	}
	return nil
}

func GetMySql() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize logger
	logger = NewLogger(p)
	return logger
}
