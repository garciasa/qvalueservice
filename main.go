package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/garciasa/qvalueservice/api"
	"github.com/garciasa/qvalueservice/api/database"
	"github.com/garciasa/qvalueservice/api/routes"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger := logrus.New()
	logger.Level = logrus.DebugLevel

	if strings.Contains(os.Getenv("ENV"), "prod") {
		logger.Formatter = new(logrus.JSONFormatter)
		logger.Level = logrus.InfoLevel
	}

	l := logger.WithField("port", "8080")

	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	qvservice := api.NewBiologyQvalueService(logger, db)
	l.Info("Binding endpoints")
	router := routes.NewRouter(qvservice)

	l.Info("Listing on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
