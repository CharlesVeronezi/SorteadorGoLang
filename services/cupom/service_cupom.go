package service_cupom

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
	connectdb "sorteador-go-lang/connectDB"
)

type CupomGeradoObj struct {
	Cupom_gerado string `json:"cupom_gerado"`
	Mensagem     string `json:"mensagem"`
}

type CupomObj struct {
	Id_cupons    int    `json:"id_cupons"`
	Codigo_cupom string `json:"codigo_cupom"`
	Id_usuario   int    `json:"id_usuario"`
	Id_sorteio   int    `json:"id_sorteio"`
}

func GerarCupom(w http.ResponseWriter, r *http.Request) {

	var cupomGeradoObj CupomGeradoObj

	RandomCrypto, _ := rand.Prime(rand.Reader, 128)

	cupomGeradoObj.Cupom_gerado = RandomCrypto.String()
	cupomGeradoObj.Mensagem = "Cupom gerado com sucesso!"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cupomGeradoObj)

}

func Cupom(w http.ResponseWriter, r *http.Request) {
	var cupomObj CupomObj
	err := json.NewDecoder(r.Body).Decode(&cupomObj)
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
	sql := `INSERT INTO cupons (codigo_cupom, id_usuario, id_sorteio) VALUES ($1, $2, $3) RETURNING id_cupons`
	err = conn.QueryRow(sql, cupomObj.Codigo_cupom, cupomObj.Id_usuario, cupomObj.Id_sorteio).Scan(&cupomObj.Id_cupons)
	if err != nil {
		http.Error(w, "Erro ao inserir usuário no banco de dados", http.StatusInternalServerError)
		return
	}

	// Retornar o usuário criado como resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cupomObj)
}
