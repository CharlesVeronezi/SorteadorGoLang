package routes

import (
	"log"
	"net/http"
	"sorteador-go-lang/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/sorteio", controllers.Sorteio)
	http.HandleFunc("/usuario", controllers.Usuario)
	http.HandleFunc("/gerar_cupom", controllers.GerarCupom)
	http.HandleFunc("/cupom", controllers.Cupom)
	http.HandleFunc("/sortear", controllers.Sortear)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
