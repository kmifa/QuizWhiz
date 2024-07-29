package entity

import (
	"time"

	"github.com/uptrace/bun"

	"github.com/kmifa/QuizWhiz/domain/model"
)

type Pet struct {
	bun.BaseModel `bun:"table:pets,alias:p"`
	ID            int64           `bun:"id,pk"`
	Name          string          `bun:"name,nullzero,notnull"`
	Status        model.PetStatus `bun:"status"`
	CreatedAt     time.Time       `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time       `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time       `bun:",soft_delete,nullzero"`
}

type PetPhoto struct {
	bun.BaseModel `bun:"table:pet_photos,alias:pp"`
	ID            int64     `bun:"id,pk"`
	PetID         int64     `bun:"pet_id"`
	PhotoURL      string    `bun:"photo_url"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt     time.Time `bun:",soft_delete,nullzero"`
}

func (p *Pet) ToModel() *model.Pet {
	return &model.Pet{
		ID:     p.ID,
		Name:   p.Name,
		Status: p.Status,
	}
}
