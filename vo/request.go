package vo

import "fmt"

type CreateCarrinhoRequest struct {
	ClienteID     string `json:"clienteId"`
	Quantidade    int64  `json:"quantidade"`
	ProdutoID     string `json:"produtoId"`
	TipoPagamento string `json:"tipo_pagamento"`
}

type UpdateCarrinhoRequest struct {
	Nome       string  `json:"nome"`
	Valor      float64 `json:"valor"`
	Quantidade int64   `json:"quantidade"`
	Descricao  string  `json:"descricao"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CreateCarrinhoRequest) Validate() error {
	if r.ClienteID == "" && r.ProdutoID == "" && r.Quantidade <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.ClienteID == "" {
		return errParamIsRequired("ClienteID", "string")
	}
	if r.ProdutoID == "" {
		return errParamIsRequired("ProdutoID", "string")
	}
	if r.Quantidade <= 0 {
		return errParamIsRequired("Quantidade", "string")
	}

	return nil
}
