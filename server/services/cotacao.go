package services

import (
	"client-server-api/config"
	"client-server-api/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func BuscaCotacao() (models.Data, error) {
	c := http.Client{Timeout: time.Millisecond * 200}
	var data models.Data
	resp, err := c.Get(config.AWESOMEAPI_URL)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return data, err
	}

	return data, nil
}
