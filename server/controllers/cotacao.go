package controllers

import (
	"client-server-api/repository"
	"client-server-api/services"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func BuscarCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	log.Println("Buscando Cotação....")
	data, error := services.BuscaCotacao()
	log.Println("Retornando Cotação....")
	if error != nil {
		http.Error(w, "Requisicao: erro ao buscar cotação", http.StatusInternalServerError)
		return
	}

	err := repository.InsertCotacao(data.USDBRL)

	if err != nil {
		http.Error(w, "Requisicao: erro ao Salvar no banco", http.StatusInternalServerError)
		return
	}

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.USDBRL)
	case <-ctx.Done():
		log.Println("Requisicao cancelada pelo cliente")
		http.Error(w, "Requisicao cancelada pelo cliente", http.StatusRequestTimeout)
	}
}

func ListarCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := repository.ListCotacao()

	if err != nil {
		panic(err)
	}

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")
	case <-ctx.Done():
		log.Println("Requisicao cancelada pelo cliente")
		http.Error(w, "Requisicao cancelada pelo cliente", http.StatusRequestTimeout)
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)

}
