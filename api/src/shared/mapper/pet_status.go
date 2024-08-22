package mapper

import (
	"github.com/kmifa/QuizWhiz/domain/model"
	"github.com/kmifa/QuizWhiz/ogen"
)

func OgenToModelPetStatus(status ogen.PetStatus) model.PetStatus {
	switch status {
	case ogen.PetStatusPending:
		return model.PetStatusPending
	case ogen.PetStatusSold:
		return model.PetStatusSold
	default:
		return model.PetStatusAvailable
	}
}

func ModelToOgenPetStatus(status model.PetStatus) ogen.PetStatus {
	switch status {
	case model.PetStatusPending:
		return ogen.PetStatusPending
	case model.PetStatusSold:
		return ogen.PetStatusSold
	default:
		return ogen.PetStatusAvailable
	}
}
