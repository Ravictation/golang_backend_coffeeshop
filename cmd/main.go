package main

import (
	"log"

	"github.com/Ravictation/golang_backend_coffeeshop/internal/pkg"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/routers"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database, err := pkg.Pgdb()
	if err != nil {
		log.Fatal(err)
	}
	router := routers.New(database)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
