package main

import (
	"database/sql"
	"net/url"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
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

	defer db.Close()
}
