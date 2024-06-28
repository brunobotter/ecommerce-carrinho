package scheamas

import (
	"gorm.io/gorm"
)

type Carrinho struct {
	gorm.Model
	NomeProduto      string
	Valor            float64
	Quantidade       int64
	DescricaoProduto string
	NomeCliente      string
	EmailCliente     string
}

type CarrinhoResponse struct {
	Id               uint    `json:"id"`
	NomeProduto      string  `json:"nome_produto"`
	Valor            float64 `json:"valor"`
	Quantidade       int64   `json:"quantidade"`
	DescricaoProduto string  `json:"descricao_produto"`
	NomeCliente      string  `json:"nome_cliente"`
	EmailCliente     string  `json:"email_cliente"`
}

func ToCarrinhoResponse(carrinho Carrinho) CarrinhoResponse {
	return CarrinhoResponse{
		Id:               carrinho.ID,
		NomeProduto:      carrinho.NomeProduto,
		Valor:            carrinho.Valor,
		Quantidade:       carrinho.Quantidade,
		DescricaoProduto: carrinho.DescricaoProduto,
		NomeCliente:      carrinho.NomeCliente,
		EmailCliente:     carrinho.EmailCliente,
	}
}
