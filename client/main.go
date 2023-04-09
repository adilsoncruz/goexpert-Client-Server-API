package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	BID string `json:"bid"`
}

func main() {
	var file *os.File
	c := http.Client{Timeout: time.Second * 10}
	resp, err := c.Get("http://localhost:8080/cotacao")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var data Cotacao
	err = json.Unmarshal(body, &data)
	file, err = os.OpenFile("cotacao.txt", os.O_APPEND, 0600)
	if err != nil {
		file, err = os.Create("cotacao.txt")
	}
	defer file.Close()
	_, err = file.WriteString("DOLAR: " + data.BID + "\n")

}
