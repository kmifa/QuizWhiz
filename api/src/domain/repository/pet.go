package repository

import (
	"context"

	"github.com/kmifa/QuizWhiz/domain/model"
)

type PetRepository interface {
	CreatePet(ctx context.Context, pet *model.Pet) (*model.Pet, error)
	UpdatePet(ctx context.Context, pet *model.Pet) (*model.Pet, error)
	DeletePet(ctx context.Context, petID int64) error
	GetPet(ctx context.Context, petID int64) (*model.Pet, error)
}
