package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/smitajshetty/go-scheduler/pkg/router"
)

func main() {
	setUpDB()

	listenAddress := "localhost:8090"
	router.Start(listenAddress)
	fmt.Println("Server listening on: ", listenAddress)
	select {}
}

func setUpDB() {
	dsn := url.URL{
		User:     url.UserPassword(os.Getenv("PG_USER"), ""),
		Scheme:   "postgres",
		Host:     "0.0.0.0:5432",
		Path:     "a",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("error")
		panic("failed to connect db")
	}

	log.WithFields(log.Fields{
		"test": db,
	}).Info("test")

	defer db.Close()
}
