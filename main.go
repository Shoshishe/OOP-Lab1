package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/rs/cors"
	"log"
	"main/controllers"
	"main/repository/postgres"
	"main/service"
	"net/http"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Couldn't init config: %s", err.Error())
	}
	db, err := postgres.NewPostgresDb(postgres.DbConfig{
		Host:     viper.GetString("db.Host"),
		Port:     viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		DbName:   viper.GetString("db.DBname"),
		Password: viper.GetString("db.Password"),
		SSLMode:  viper.GetString("db.SSLmode"),
	})
	if err != nil {
		log.Fatalf("Failed to init db %s", err.Error())
	}
	repos := service.NewRepository(
		postgres.NewAuthPostgres(db),
		postgres.NewBankPostgres(db),
		postgres.NewBankAccountPostgres(db),
	)
	serv := service.NewService(*repos)
	controllers := controllers.NewController(serv)

	mux := http.NewServeMux()
	muxHandler := cors.Default().Handler(mux)
	controllers.RegisterRoutes(mux)
	err = http.ListenAndServe(viper.GetString("net.Host")+viper.GetString("net.Port"), muxHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
