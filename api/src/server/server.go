package server

import (
	"github.com/kmifa/QuizWhiz/ogen"
	"github.com/kmifa/QuizWhiz/usecase"
)

type quizWhizServer struct {
	petUsecase usecase.PetUsecase
}

func NewQuizWhizServer(petUsecase usecase.PetUsecase) ogen.Handler {
	return &quizWhizServer{
		petUsecase: petUsecase,
	}
}
