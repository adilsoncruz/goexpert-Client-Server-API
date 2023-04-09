package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *sql.DB

func Connect() (*sql.DB, error) {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			log.Println("Creating Connection instance now.")
			db, err := sql.Open("sqlite3", "./cotacao.db")

			if err != nil {
				panic(err)
			}
			singleInstance = db
			log.Println("Connection instance created.")
		} else {
			log.Println("Connection instance already created.")
		}
	} else {
		log.Println("Connection instance already created.")
	}

	return singleInstance, nil
}
