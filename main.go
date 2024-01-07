package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/zakisk/socio-server/data"
	"github.com/zakisk/socio-server/handlers"
	"github.com/zakisk/socio-server/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Unable to load .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal().Msg("Unable to open database conenction")
	}

	dbHandler := data.NewDBHandler(db)
	handlerInstance := handlers.NewHandlerInstance(log, dbHandler)

	router := router.NewRouter(handlerInstance)

	go func() {
		err = router.S.Run(":3001")
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Info().Str("signal", sig.String()).Msg("Got signal")
}
