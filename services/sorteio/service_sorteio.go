package service_sorteio

import (
	"encoding/json"
	"net/http"
	connectdb "sorteador-go-lang/connectDB"
)

type Sorteio struct {
	Id_sorteio   int    `json:"id_sorteio"`
	Nome_sorteio string `json:"nome_sorteio"`
	Data_sorteio string `json:"data_sorteio"`
}

func CriarSorteio(w http.ResponseWriter, r *http.Request) {
	var sorteio Sorteio
	err := json.NewDecoder(r.Body).Decode(&sorteio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Conexão com o banco de dados
	conn, err := connectdb.Connection()
	if err != nil {
		http.Error(w, "Erro ao conectar no banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Inserir no banco de dados
	sql := `INSERT INTO sorteio (nome_sorteio, data_sorteio) VALUES ($1, $2) RETURNING id_sorteio`
	err = conn.QueryRow(sql, sorteio.Nome_sorteio, sorteio.Data_sorteio).Scan(&sorteio.Id_sorteio)
	if err != nil {
		http.Error(w, "Erro ao inserir usuário no banco de dados", http.StatusInternalServerError)
		return
	}

	// Retornar o usuário criado como resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sorteio)
}
