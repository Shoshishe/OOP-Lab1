package main

import (
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/spf13/viper"
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
		postgres.NewReverserRepository(db),
	)
	serv := service.NewService(*repos)
	controllers := controllers.NewController(serv)

	mux := http.NewServeMux()
	muxHandler := cors.New(
		cors.Options{
			AllowedOrigins: []string{"http://localhost:5173"},
			AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
			AllowedMethods: []string{"GET", "POST", "PUT", "UPDATE", "DELETE", "OPTIONS"},
		},
	).Handler(mux)
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
