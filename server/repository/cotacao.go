package repository

import (
	"client-server-api/lib/db"
	"client-server-api/models"
	"context"
	"log"
	"time"
)

func InsertCotacao(c models.Cotacao) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	log.Println("inserir", c, c.Code)
	connection, err := db.Connect()
	if err != nil {
		return err
	}
	stmt, err := connection.Prepare("insert into cotacao(Code, CodeIn, NamePrice, Hight, Low, PctChange, BID, Ask, TimestampPrice, CreateDate) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, c.Code, c.CodeIn, c.Name, c.Hight, c.Low, c.PctChange, c.BID, c.Ask, c.TimestampPrice, c.CreateDate)

	if err != nil {
		return err
	}

	log.Println("Cotacao salva")
	return nil
}

func ListCotacao() ([]models.Cotacao, error) {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	connection, err := db.Connect()
	if err != nil {
		return nil, err
	}
	var data []models.Cotacao
	log.Println("consultar no banco")
	rows, err := connection.Query(`SELECT Code, CodeIn, NamePrice, Hight, Low, PctChange, BID, Ask, TimestampPrice, CreateDate FROM cotacao`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var cotacao models.Cotacao

		if err := rows.Scan(&cotacao.Code, &cotacao.CodeIn, &cotacao.NamePrice, &cotacao.Hight, &cotacao.Low, &cotacao.PctChange, &cotacao.BID, &cotacao.Ask, &cotacao.TimestampPrice, &cotacao.CreateDate); err != nil {
			log.Println(err.Error())
		}

		data = append(data, cotacao)
	}

	log.Println("Cotacao salva")
	return data, nil
}
