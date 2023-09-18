package main

import (
	"context"
	"eff_mob_test"
	"eff_mob_test/pkg/handler"
	"eff_mob_test/pkg/repository"
	"eff_mob_test/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {

	initEnv()

	ctx := context.Background()

	logrus.Infoln("Connecting to a DB")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infoln("Connection to the DB is established")

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(eff_mob_test.Server)

	logrus.Println("Starting server")

	go func() {
		if err := srv.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Infoln("Initiating Kafka processing")

	handlers.InitKafka(ctx)

}

func initEnv() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}
}
