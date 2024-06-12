package service_usuario

import (
	"encoding/json"
	"net/http"
	connectdb "sorteador-go-lang/connectDB"
	"time"
)

type Usuario struct {
	Id_usuario           int    `json:"id_usuario"`
	Nome_usuario         string `json:"nome_usuario"`
	Data_criacao_usuario string `json:"data_criacao_usuario"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
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

	// Definir data de criação atual se não estiver definida
	if usuario.Data_criacao_usuario == "" {
		usuario.Data_criacao_usuario = time.Now().Format("2006-01-02 15:04:05")
	}

	// Inserir no banco de dados
	sql := `INSERT INTO usuario (nome_usuario, data_criacao_usuario) VALUES ($1, $2) RETURNING id_usuario`
	err = conn.QueryRow(sql, usuario.Nome_usuario, usuario.Data_criacao_usuario).Scan(&usuario.Id_usuario)
	if err != nil {
		http.Error(w, "Erro ao inserir usuário no banco de dados", http.StatusInternalServerError)
		return
	}

	// Retornar o usuário criado como resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}
