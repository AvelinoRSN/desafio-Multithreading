package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BrasilAPIResponse struct {
	CEP        		string `json:"cep"`
	Street 	   		string `json:"street"`	
	Neighborhood    string `json:"neighborhood"`
	City     		string `json:"city"`	
	State     		string `json:"state"`	
}

type ViaCEPResponse struct {
	CEP        		string `json:"cep"`
	Logradouro 		string `json:"logradouro"`
	Bairro     		string `json:"bairro"`
	Cidade     		string `json:"localidade"`
	Estado     		string `json:"uf"`
}

type Endereco struct {
	CEP        		string
	Logradouro 		string
	Bairro     		string
	Cidade     		string
	Estado     		string
	Origem     		string
}

func main() {
	cep:= "13175658"
	timeout := time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan Endereco,2)

	go BuscarBrasilAPI(ctx, cep, ch)
	go BuscarViaCEP(ctx, cep, ch)

	select {
	case endereco := <-ch:
		fmt.Printf("Resposta mais rápida da API: %s\nCEP: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nEstado: %s\n", 
		endereco.Origem, endereco.CEP, endereco.Logradouro, endereco.Bairro, endereco.Cidade, endereco.Estado)
	
	case <-ctx.Done():
		fmt.Println("Erro: Tempo excedido ao buscar o CEP")
	}
}

func BuscarBrasilAPI(ctx context.Context, cep string, ch chan<- Endereco) {
	url := "https://brasilapi.com.br/api/cep/v1/"+ cep 
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {	
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	var apiResp BrasilAPIResponse
	err = json.NewDecoder(res.Body).Decode(&apiResp)
	if err != nil {
		return
	}
	end := Endereco{
		CEP:        apiResp.CEP,
		Logradouro: apiResp.Street,
		Bairro:     apiResp.Neighborhood,
		Cidade:     apiResp.City,
		Estado:     apiResp.State,	
		Origem:     "BrasilAPI",
	}
	ch <- end
}

func BuscarViaCEP(ctx context.Context, cep string, ch chan<- Endereco){
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	var apiResp ViaCEPResponse
	err = json.NewDecoder(res.Body).Decode(&apiResp)
	if err != nil {
		return
	}
	end := Endereco{
		CEP:        apiResp.CEP,
		Logradouro: apiResp.Logradouro,
		Bairro:     apiResp.Bairro,
		Cidade:     apiResp.Cidade,
		Estado:     apiResp.Estado,	
		Origem:     "ViaCEP",
	}
	ch <- end
}