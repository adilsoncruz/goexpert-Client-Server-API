package main

import (
	"client-server-api/controllers"
	"net/http"
)

type ErrorMessage struct {
	Code    int
	Message string
}

func main() {
	http.HandleFunc("/cotacao", controllers.BuscarCotacaoHandler)
	http.HandleFunc("/history", controllers.ListarCotacaoHandler)
	http.ListenAndServe(":8080", nil)

}
