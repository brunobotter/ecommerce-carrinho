package service

import (
	"encoding/json"

	"github.com/brunobotter/ecommerce-carrinho/configs"
	"github.com/brunobotter/ecommerce-carrinho/integration"
	"github.com/brunobotter/ecommerce-carrinho/scheamas"
	"github.com/brunobotter/ecommerce-carrinho/vo"
	"gorm.io/gorm"
)

// db é a instância do banco de dados global
var db *gorm.DB

func InitializeService(database *gorm.DB) {
	db = database
}
func AdicionarAoCarrinho(request vo.CreateCarrinhoRequest) (scheamas.CarrinhoResponse, error) {
	logger := configs.GetLogger("AdicionarAoCarrinho")

	cliente, err := integration.GetCliente(request.ClienteID)
	if err != nil {
		return scheamas.CarrinhoResponse{}, err
	}
	logger.Debug("api cliente ok")

	venda, err := integration.PostVenda(request)
	if err != nil {
		return scheamas.CarrinhoResponse{}, err
	}
	logger.Debugf("api venda ok ")

	carrinho := scheamas.Carrinho{
		Quantidade:       request.Quantidade,
		NomeProduto:      venda.Nome,
		Valor:            venda.Valor,
		DescricaoProduto: venda.Descricao,
		NomeCliente:      cliente.Nome,
		EmailCliente:     cliente.Email,
	}

	if err := db.Create(&carrinho).Error; err != nil {
		logger.Errorf("erro ao salvar no banco %v", err)
		return scheamas.CarrinhoResponse{}, err
	}
	carrinhoResponse := scheamas.ToCarrinhoResponse(carrinho)
	logger.Debugf("salvou no banco ")
	// Preparar dados de pagamento para envio à fila SQS
	pagamentoData := struct {
		CarrinhoID       uint    `json:"carrinho_id"`
		NomeCliente      string  `json:"nome_cliente"`
		EmailCliente     string  `json:"email_cliente"`
		NomeProduto      string  `json:"nome_produto"`
		Valor            float64 `json:"valor"`
		Quantidade       int64   `json:"quantidade"`
		DescricaoProduto string  `json:"descricao_produto"`
		TipoPagamento    string  `json:"tipoPagamento"`
	}{
		CarrinhoID:       carrinho.ID,
		NomeCliente:      cliente.Nome,
		EmailCliente:     cliente.Email,
		NomeProduto:      venda.Nome,
		Valor:            venda.Valor,
		Quantidade:       request.Quantidade,
		DescricaoProduto: venda.Descricao,
		TipoPagamento:    request.TipoPagamento,
	}

	pagamentoJSON, err := json.Marshal(pagamentoData)
	if err != nil {
		logger.Errorf("erro na deserealização %v", err)
		return scheamas.CarrinhoResponse{}, err
	}

	// Enviar para fila de pagamento
	queueURL := "https://sqs.us-east-1.amazonaws.com/730335442778/Pagamento.fifo" // Substitua pela URL da sua fila SQS
	err = integration.SendMessageToSQS(queueURL, string(pagamentoJSON))
	if err != nil {
		logger.Errorf("erro envio para fila %v", err)
		return scheamas.CarrinhoResponse{}, err
	}

	return carrinhoResponse, nil
}

func ShowCarrinho(id string) (scheamas.Carrinho, error) {
	carrinho := scheamas.Carrinho{}
	if err := db.First(&carrinho, id).Error; err != nil {
		return carrinho, err
	}

	return carrinho, nil
}

func ListCarrinhos() ([]scheamas.Carrinho, error) {
	var carrinhos []scheamas.Carrinho
	if err := db.Find(&carrinhos).Error; err != nil {
		return nil, err
	}

	return carrinhos, nil
}
