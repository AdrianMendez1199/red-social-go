package main

import (
	"github.com/AdrianMendez1199/red-social-go/handlers"
	"github.com/AdrianMendez1199/red-social-go/db"
)

func main() {
	handlers.Handler()
	db.ConnectDB()
}