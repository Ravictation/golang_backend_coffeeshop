package main

import (
	"log"

	"github.com/Ravictation/golang_backend_coffeeshop/internal/pkg"
	"github.com/Ravictation/golang_backend_coffeeshop/internal/routers"
	"github.com/asaskevich/govalidator"
	"github.com/spf13/viper"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	viper.SetConfigName("env.dev")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

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
