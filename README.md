# Aplicação que realiza a busca de CEP

Este projeto em Go realiza consultas simultâneas a duas APIs diferentes de CEP e retorna o resultado da API que responder mais rápido.

## APIs Utilizadas

- **BrasilAPI**: `https://brasilapi.com.br/api/cep/v1/{CEP}`
- **ViaCEP**: `http://viacep.com.br/ws/{CEP}/json/`

## Requisitos

- O programa deve realizar as requisições em paralelo para ambas as APIs.
- Apenas a resposta mais rápida deve ser considerada, descartando a outra.
- O tempo máximo de resposta deve ser de **1 segundo**. Caso contrário, um erro de timeout será exibido.
- Os resultados devem ser exibidos no terminal, incluindo qual API forneceu a resposta.
- Como os campos das APIs diferem, a resposta é padronizada.

## Estrutura do Código

- `BuscarBrasilAPI()`: Realiza a requisição na BrasilAPI.
- `BuscarViaCEP`: Realiza a requisição no ViaCEP.
- `main()`: Controla a execução do programa e exibe os resultados.

## Como Executar

1. Instale o Go (caso ainda não tenha instalado): [https://go.dev/dl/](https://go.dev/dl/)
2. Clone este repositório:
   ```sh
   git clone (https://github.com/AvelinoRSN/desafio-Multithreading.git)
   ```
3. Navegue até o diretório do projeto:
   ```sh
   cd seu-repositorio
   ```
4. Execute o programa:
   ```sh
   go run main.go
   ```

## Exemplo de Saída

```sh
Resposta mais rápida da API ViaCEP:
CEP: 01153000
Logradouro: Rua Exemplo
Bairro: Bairro Teste
Cidade: São Paulo
Estado: SP
```

Caso o tempo limite seja atingido:
```sh
Erro: Timeout ao buscar o CEP
```

## Autor

Avelino

