package handler

import (
	"github.com/brunobotter/ecommerce-carrinho/configs"
	"github.com/brunobotter/ecommerce-carrinho/service"
	"gorm.io/gorm"
)

var (
	logger *configs.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = configs.GetLogger("handler")
	db = configs.GetMySql()
	service.InitializeService(db)
}
