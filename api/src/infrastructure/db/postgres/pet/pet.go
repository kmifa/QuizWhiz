package pet

import (
	"context"
	"database/sql"
	"time"

	"github.com/kmifa/QuizWhiz/domain/model"
	"github.com/kmifa/QuizWhiz/domain/repository"
	"github.com/kmifa/QuizWhiz/infrastructure/db/postgres"
	"github.com/kmifa/QuizWhiz/infrastructure/db/postgres/entity"
)

func NewPetRepository(db *postgres.DB) repository.PetRepository {
	return &petRepository{
		db: db,
	}
}

type petRepository struct {
	db *postgres.DB
}

func (r *petRepository) CreatePet(ctx context.Context, pet *model.Pet) (*model.Pet, error) {

	now := time.Now()

	p := &entity.Pet{
		ID:        pet.ID,
		Name:      pet.Name,
		Status:    pet.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if _, err := r.db.GetDB(ctx).
		NewInsert().
		Model(p).
		Returning("id").
		Exec(ctx); err != nil {
		return nil, err
	}
	return p.ToModel(), nil
}

func (r *petRepository) UpdatePet(ctx context.Context, pet *model.Pet) (*model.Pet, error) {
	now := time.Now()
	p := &entity.Pet{
		ID:        pet.ID,
		Name:      pet.Name,
		Status:    pet.Status,
		UpdatedAt: now,
	}

	res, err := r.db.GetDB(ctx).
		NewUpdate().
		Model(p).
		Returning("*").
		WherePK().
		Exec(ctx)
	ra, err := res.RowsAffected()
	if err != nil {
		return nil, err
	} else if ra == 0 {
		return nil, err
	}

	return p.ToModel(), nil
}

func (r *petRepository) DeletePet(ctx context.Context, petID int64) error {
	res, err := r.db.GetDB(ctx).
		NewDelete().
		Model(&entity.Pet{}).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return err
	} else if ra == 0 {
		return err
	}

	return nil
}

func (r *petRepository) GetPet(ctx context.Context, petID int64) (*model.Pet, error) {
	var p entity.Pet
	if err := r.db.GetDB(ctx).
		NewSelect().
		NewSelect().
		Model(&p).
		WherePK().
		Scan(ctx); err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return p.ToModel(), nil
}
