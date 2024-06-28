package handler

import (
	"github.com/brunobotter/ecommerce-carrinho/config"
	"github.com/brunobotter/ecommerce-carrinho/service"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetMySql()
	service.InitializeService(db)
}
