package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	// Cria um contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Cria a requisição com o contexto
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	// Executa a requisição
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta: %v", err)
	}

	// Verifica o status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Erro na requisição: %s", string(body))
	}

	// Decodifica a resposta
	var cotacao Cotacao
	if err := json.Unmarshal(body, &cotacao); err != nil {
		log.Fatalf("Erro ao decodificar resposta: %v", err)
	}

	// Cria ou abre o arquivo para escrita
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatalf("Erro ao criar arquivo: %v", err)
	}
	defer file.Close()

	// Escreve a cotação no arquivo
	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", cotacao.Bid))
	if err != nil {
		log.Fatalf("Erro ao escrever no arquivo: %v", err)
	}

	fmt.Printf("Cotação do dólar salva no arquivo cotacao.txt: %s\n", cotacao.Bid)
}
