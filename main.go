package main

import (
	"log"
	"main/infrastructure"
	"main/infrastructure/postgres"
	 _"github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Couldn't init config: %s", err.Error())
	}
	db, err := postgres.NewPostgresDb(postgres.DbConfig{
		Host: viper.GetString("dn.Host"),
		Port: viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		DbName: viper.GetString("db.DBname"),
		SSLMode: viper.GetString("db.SSLmode"),
	})
	if err != nil {
		log.Fatalf("Failed to init db %s", err.Error())
	}
	repos := infrastructure.NewRepository(
		postgres.NewAuthPostgres(db),
	)
	//Just to avoid unused import
	repos.GetUser("","")
}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
