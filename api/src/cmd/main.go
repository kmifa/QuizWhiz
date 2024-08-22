package main

import (
	"log"
	"net/http"

	"github.com/kmifa/QuizWhiz/config"
	"github.com/kmifa/QuizWhiz/infrastructure/db/postgres"
	"github.com/kmifa/QuizWhiz/infrastructure/db/postgres/pet"
	"github.com/kmifa/QuizWhiz/ogen"
	"github.com/kmifa/QuizWhiz/server"
	"github.com/kmifa/QuizWhiz/usecase"
	"github.com/kmifa/QuizWhiz/utilities/errors"
)

func main() {
	if err := http.ListenAndServe(":8080", setServer()); err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}

func setServer() *ogen.Server {
	config.SetEnv()

	db := postgres.NewDB()

	petRepository := pet.NewPetRepository(db)

	petUsecase := usecase.NewPetUsecase(db, petRepository)

	s, err := ogen.NewServer(
		server.NewQuizWhizServer(petUsecase),
		ogen.WithErrorHandler(errors.ErrorHandler),
	)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	log.Printf("server started")
	return s
}
