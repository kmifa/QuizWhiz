package port

import "github.com/kmifa/QuizWhiz/ogen"

type Pet struct {
	ID     int64
	Name   string
	Status ogen.PetStatus
}

// CreatePetInputPort is the input port for creating a pet.
type CreatePetInputPort struct {
	ID   int64
	Name string
}

// CreatePetOutputPort is the output port for creating a pet.
type CreatePetOutputPort Pet

// UpdatePetInputPort is the input port for updating a pet.
type UpdatePetInputPort struct {
	ID     int64
	Name   string
	Status ogen.PetStatus
}

// UpdatePetOutputPort is the output port for updating a pet.
type UpdatePetOutputPort Pet

// GetPetOutputPort is the output port for getting a pet.
type GetPetOutputPort Pet
