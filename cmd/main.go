package main

import (
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("postgres", "scheduler")
	if err != nil {
		panic("failed to connect db")
	}

	defer db.Close()
}
