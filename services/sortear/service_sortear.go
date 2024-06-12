package service_sortear

import (
	"encoding/json"
	"math/rand"
	"net/http"
	connectdb "sorteador-go-lang/connectDB"
	"time"
)

type SortearObj struct {
	Id_sorteio int `json:"id_sorteio"`
}

type CupomObj struct {
	Id_cupons    int    `json:"id_cupons"`
	Codigo_cupom string `json:"codigo_cupom"`
	Id_usuario   int    `json:"id_usuario"`
	Id_sorteio   int    `json:"id_sorteio"`
}

func selecionaCupom(cupons []CupomObj) (resultado CupomObj) {
	if len(cupons) == 0 {
		return
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(cupons))
	return cupons[index]
}

func SortearCupom(w http.ResponseWriter, r *http.Request) {
	var sortear SortearObj
	err := json.NewDecoder(r.Body).Decode(&sortear)
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

	sql := `SELECT id_cupons, codigo_cupom, id_usuario, id_sorteio FROM cupons WHERE id_sorteio = $1`
	rows, err := conn.Query(sql, sortear.Id_sorteio)
	if err != nil {
		http.Error(w, "Erro ao buscar cupons no banco de dados", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cupons []CupomObj
	for rows.Next() {
		var cupom CupomObj
		err := rows.Scan(&cupom.Id_cupons, &cupom.Codigo_cupom, &cupom.Id_usuario, &cupom.Id_sorteio)
		if err != nil {
			http.Error(w, "Erro ao ler dados dos cupons", http.StatusInternalServerError)
			return
		}
		cupons = append(cupons, cupom)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Erro ao iterar pelos cupons", http.StatusInternalServerError)
		return
	}

	cupomPremiado := selecionaCupom(cupons)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cupomPremiado)
}
