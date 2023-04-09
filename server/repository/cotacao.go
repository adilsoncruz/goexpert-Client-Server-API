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

	connection, err := db.Connect()
	if err != nil {
		return err
	}
	stmt, err := connection.Prepare("insert into Cotacao(Code, CodeIn, NamePrice, Hight, Low, PctChange, BID, Ask, TimestampPrice, CreateDate) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, c.Code, c.CodeIn, c.Name, c.Hight, c.Low, c.PctChange, c.BID, c.Ask, c.Timestamp, c.CreateDate)

	if err != nil {
		return err
	}

	log.Println("Cotacao salva")
	return nil
}
