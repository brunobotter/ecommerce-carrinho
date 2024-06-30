package handler

import (
	"net/http"

	"github.com/brunobotter/ecommerce-carrinho/configs"
	"github.com/brunobotter/ecommerce-carrinho/service"
	"github.com/brunobotter/ecommerce-carrinho/vo"
	"github.com/gin-gonic/gin"
)

func CreateCarrinhoHandler(ctx *gin.Context) {
	logger := configs.GetLogger("handler")
	var request vo.CreateCarrinhoRequest
	if err := ctx.Bind(&request); err != nil {
		vo.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	accessKeyID := ctx.GetHeader("X-AWS-Access-Key-ID")
	secretAccessKey := ctx.GetHeader("X-AWS-Secret-Access-Key")
	logger.Debugf("ace %s", accessKeyID)
	region := ctx.GetHeader("X-AWS-Region")

	carrinho, err := service.AdicionarAoCarrinho(request, accessKeyID, secretAccessKey, region)
	if err != nil {
		vo.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	vo.SendSuccess(ctx, "Produto adicionado ao carrinho criado com sucesso", carrinho)
}
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func ShowCarrinhoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		vo.SendError(ctx, http.StatusBadRequest, "ID é obrigatório")
		return
	}

	carrinho, err := service.ShowCarrinho(id)
	if err != nil {
		vo.SendError(ctx, http.StatusNotFound, "Carrinho não encontrado")
		return
	}

	vo.SendSuccess(ctx, "Carrinho encontrado", carrinho)
}

func ListCarrinhoHandler(ctx *gin.Context) {
	carrinhos, err := service.ListCarrinhos()
	if err != nil {
		vo.SendError(ctx, http.StatusInternalServerError, "Erro ao listar carrinhos")
		return
	}

	vo.SendSuccess(ctx, "Carrinhos listados com sucesso", carrinhos)
}
