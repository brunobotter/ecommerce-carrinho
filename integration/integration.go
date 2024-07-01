package integration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/brunobotter/ecommerce-carrinho/configs"
	"github.com/brunobotter/ecommerce-carrinho/vo"
)

const (
	clienteAPIURL  = "http://ecommerce-usuario-alb-1247437976.us-east-1.elb.amazonaws.com:8080/api/v1/usuario"     // URL da API de cliente
	vendaAPIURL    = "http://ecommerce-produto-alb-54842778.us-east-1.elb.amazonaws.com:8081/api/v1/produto/venda" // URL da API de venda
	requestTimeout = 15 * time.Second
)

type ClienteResponseData struct {
	ID        uint   `json:"id"`
	Nome      string `json:"nome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
	Documento string `json:"documento"`
	Endereco  string `json:"endereco"`
}

type ClienteResponse struct {
	Data ClienteResponseData `json:"data"`
}

type VendaResponseData struct {
	Id         uint    `json:"Id"`
	Nome       string  `json:"Nome"`
	Quantidade int64   `json:"Quantidade"`
	Valor      float64 `json:"Valor"`
	Descricao  string  `json:"Descricao"`
}

type VendaResponse struct {
	Data VendaResponseData `json:"data"`
}

type VendaRequest struct {
	ProdutoID  string `json:"produtoId"`
	Quantidade int64  `json:"quantidade"`
}

func GetCliente(clienteID string) (ClienteResponseData, error) {
	url := fmt.Sprintf("%s/%s", clienteAPIURL, clienteID)
	client := http.Client{Timeout: requestTimeout}

	resp, err := client.Get(url)
	if err != nil {
		return ClienteResponseData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ClienteResponseData{}, errors.New("failed to get usuario data")
	}

	var cliente ClienteResponse
	err = json.NewDecoder(resp.Body).Decode(&cliente)
	if err != nil {
		return ClienteResponseData{}, err
	}

	return cliente.Data, nil
}

func PostVenda(request vo.CreateCarrinhoRequest) (VendaResponseData, error) {
	logger := configs.GetLogger("produto")
	url := vendaAPIURL
	// Converter request para JSON
	vendaRequest := VendaRequest{
		ProdutoID:  request.ProdutoID,
		Quantidade: request.Quantidade,
	}
	jsonData, err := json.Marshal(vendaRequest)
	if err != nil {
		return VendaResponseData{}, err
	}

	// Criar a requisição PATCH
	client := http.Client{Timeout: requestTimeout}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return VendaResponseData{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Enviar a requisição PATCH
	resp, err := client.Do(req)
	if err != nil {
		return VendaResponseData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return VendaResponseData{}, err
	}

	var venda VendaResponse
	if err := json.NewDecoder(resp.Body).Decode(&venda); err != nil {
		return VendaResponseData{}, err
	}

	logger.Debugf("Decoded venda response: %+v", venda.Data)

	return venda.Data, nil
}
