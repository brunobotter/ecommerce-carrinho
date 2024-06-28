package config

import (
	"github.com/brunobotter/ecommerce-carrinho/scheamas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeSql() (*gorm.DB, error) {
	logger := GetLogger("mysql")
	dsn := "admin:bruno171191@tcp(produto.cfwgqgmc8vat.us-east-1.rds.amazonaws.com:3306)/produto?charset=utf8mb4&parseTime=true&loc=Local"
	//create db and connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("mysql carrinho error: %v", err)
		return nil, err
	}
	//migrate scheama
	err = db.AutoMigrate(&scheamas.Carrinho{})
	if err != nil {
		logger.Errorf("mysql automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
