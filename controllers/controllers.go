package controllers

import (
	"fmt"
	"net/http"
	service_cupom "sorteador-go-lang/services/cupom"
	service_sortear "sorteador-go-lang/services/sortear"
	service_sorteio "sorteador-go-lang/services/sorteio"
	service_usuario "sorteador-go-lang/services/usuario"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Usuario(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		service_usuario.CriarUsuario(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}

func Sorteio(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		service_sorteio.CriarSorteio(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}

func GerarCupom(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		service_cupom.GerarCupom(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}

func Cupom(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		service_cupom.Cupom(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}

func Sortear(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		service_sortear.SortearCupom(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}
