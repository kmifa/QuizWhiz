package server

import (
	"context"
	"log"

	"github.com/morikuni/failure"

	"github.com/kmifa/QuizWhiz/ogen"
	"github.com/kmifa/QuizWhiz/utilities/errors"
)

// POST /pet
func (s *quizWhizServer) AddPet(ctx context.Context, req *ogen.Pet) (*ogen.Pet, error) {
	return nil, failure.New(errors.UnImplemented, failure.Message("not implemented"))
}

// DELETE /pet/{petId}
func (s *quizWhizServer) DeletePet(ctx context.Context, params ogen.DeletePetParams) error {
	return failure.New(errors.UnImplemented, failure.Message("not implemented"))
}

// GET /pet/{petId}
func (s *quizWhizServer) GetPetById(ctx context.Context, params ogen.GetPetByIdParams) (ogen.GetPetByIdRes, error) {
	log.Printf("GetPetById: %v", params.PetId)
	return nil, failure.New(errors.UnImplemented, failure.Message("not implemented"))
}

// POST /pet/{petId}
func (s *quizWhizServer) UpdatePet(ctx context.Context, params ogen.UpdatePetParams) error {
	return failure.New(errors.UnImplemented, failure.Message("not implemented"))
}
