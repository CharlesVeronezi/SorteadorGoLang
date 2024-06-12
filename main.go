package main

import (
	"fmt"
	"sorteador-go-lang/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
