package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/AdrianMendez1199/red-social-go/middlewares"
	"github.com/AdrianMendez1199/red-social-go/routes"
)

/* Handler list http Server  */
func Handler() {
	router := mux.NewRouter()

	router.HandlerFunc("/register", middlewares.CheckDB()).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}