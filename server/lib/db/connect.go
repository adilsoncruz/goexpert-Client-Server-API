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
			db, err := sql.Open("sqlite3", "db")

			if err != nil {
				panic(err)
			}
			db.Exec(`CREATE TABLE cotacao ( 
				uid INTEGER PRIMARY KEY AUTOINCREMENT,   
				 Code       varchar(255) NULL,
				CodeIn     varchar(255) NULL,
				NamePrice       varchar(255) NULL,
				Hight      varchar(255) NULL,
				Low      varchar(255) NULL,
				PctChange  varchar(255) NULL,
				BID        varchar(255) NULL,
				Ask        varchar(255) NULL,
				TimestampPrice  varchar(255) NULL,
				CreateDate varchar(255) NULL
				);`)
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
