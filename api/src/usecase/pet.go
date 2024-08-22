package usecase

import (
	"context"

	"github.com/kmifa/QuizWhiz/domain/model"
	"github.com/kmifa/QuizWhiz/domain/repository"
	"github.com/kmifa/QuizWhiz/infrastructure/db/postgres"
	"github.com/kmifa/QuizWhiz/shared/mapper"
	"github.com/kmifa/QuizWhiz/usecase/port"
)

type PetUsecase interface {
	CreatePet(ctx context.Context, input *port.CreatePetInputPort) (*port.CreatePetOutputPort, error)
	DeletePet(ctx context.Context, ID int64) error
	GetPet(ctx context.Context, ID int64) (*port.GetPetOutputPort, error)
	UpdatePet(ctx context.Context, input *port.UpdatePetInputPort) (*port.UpdatePetOutputPort, error)
}

type petUsecase struct {
	db            *postgres.DB
	petRepository repository.PetRepository
}

func NewPetUsecase(db *postgres.DB, petRepository repository.PetRepository) PetUsecase {
	return &petUsecase{
		db:            db,
		petRepository: petRepository,
	}
}

func (u *petUsecase) CreatePet(ctx context.Context, input *port.CreatePetInputPort) (*port.CreatePetOutputPort, error) {
	res, err := u.petRepository.CreatePet(ctx, &model.Pet{
		ID:     input.ID,
		Name:   input.Name,
		Status: model.PetStatusAvailable,
	})
	if err != nil {
		return nil, err
	}

	return &port.CreatePetOutputPort{
		ID:     res.ID,
		Name:   res.Name,
		Status: mapper.ModelToOgenPetStatus(res.Status),
	}, nil
}

func (u *petUsecase) DeletePet(ctx context.Context, ID int64) error {
	return u.petRepository.DeletePet(ctx, ID)
}

func (u *petUsecase) GetPet(ctx context.Context, ID int64) (*port.GetPetOutputPort, error) {
	res, err := u.petRepository.GetPet(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &port.GetPetOutputPort{
		ID:     res.ID,
		Name:   res.Name,
		Status: mapper.ModelToOgenPetStatus(res.Status),
	}, nil
}

func (u *petUsecase) UpdatePet(ctx context.Context, input *port.UpdatePetInputPort) (*port.UpdatePetOutputPort, error) {
	res, err := u.petRepository.UpdatePet(ctx, &model.Pet{
		ID:     input.ID,
		Name:   input.Name,
		Status: mapper.OgenToModelPetStatus(input.Status),
	})
	if err != nil {
		return nil, err
	}

	return &port.UpdatePetOutputPort{
		ID:     res.ID,
		Name:   res.Name,
		Status: mapper.ModelToOgenPetStatus(res.Status),
	}, nil
}
