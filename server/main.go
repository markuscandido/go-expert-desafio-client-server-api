package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

type APIResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cotacoes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			valor TEXT NOT NULL,
			data DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctxAPI, cancelAPI := context.WithTimeout(r.Context(), 200*time.Millisecond)
		defer cancelAPI()

		req, err := http.NewRequestWithContext(ctxAPI, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Erro ao criar requisição: %v", err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Erro ao buscar cotação: "+err.Error(), http.StatusInternalServerError)
			log.Printf("Erro ao buscar cotação: %v", err)
			return
		}
		defer resp.Body.Close()

		var apiResponse APIResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			http.Error(w, "Erro ao decodificar resposta: "+err.Error(), http.StatusInternalServerError)
			log.Printf("Erro ao decodificar resposta: %v", err)
			return
		}

		ctxDB, cancelDB := context.WithTimeout(r.Context(), 10*time.Millisecond)
		defer cancelDB()

		_, err = db.ExecContext(ctxDB, "INSERT INTO cotacoes (valor) VALUES (?)", apiResponse.USDBRL.Bid)
		if err != nil {
			http.Error(w, "Erro ao salvar no banco de dados: "+err.Error(), http.StatusInternalServerError)
			log.Printf("Erro ao salvar no banco de dados: %v", err)
			return
		}

		cotacao := Cotacao{Bid: apiResponse.USDBRL.Bid}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cotacao)
	})

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
