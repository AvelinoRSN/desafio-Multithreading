package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Endereco struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`	
	Bairro     string `json:"bairro"`
	Cidade     string `json:"localidade"`	
	Estado     string `json:"uf"`	
	Origem     string
}

func main() {
	cep:= "13175658"
	timeout := time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan Endereco,2)

	go BuscarViaCEP(ctx, cep, ch)
	go BuscarBrasilAPI(ctx, cep, ch)

	select {
	case endereco := <-ch:
		fmt.Printf("Resposta mais rápida da API: %s\nCEP: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nEstado: %s\n", 
		endereco.Origem, endereco.CEP, endereco.Logradouro, endereco.Bairro, endereco.Cidade, endereco.Estado)
	
	case <-ctx.Done():
		fmt.Println("Erro: Tempo excedido ao buscar o CEP")
	}
}

func BuscarBrasilAPI(ctx context.Context, cep string, ch chan<- Endereco) {
	url := "https://brasilapi.com.br/api/cep/v1/"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {	
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	var endereco Endereco
	err = json.NewDecoder(res.Body).Decode(&endereco)
	if err != nil {
		return
	}
	endereco.Origem = "BrasilAPI"
	ch <- endereco
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

	var endereco Endereco
	err = json.NewDecoder(res.Body).Decode(&endereco)
	if err != nil {
		return
	}
	endereco.Origem = "ViaCEP"
	ch <- endereco
}