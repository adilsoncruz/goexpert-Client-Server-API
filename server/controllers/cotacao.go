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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	repository.InsertCotacao(data.USDBRL)

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")
	case <-ctx.Done():
		log.Println("Requisicao cancelada pelo cliente")
		http.Error(w, "Requisicao cancelada pelo cliente", http.StatusRequestTimeout)
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.USDBRL)

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
