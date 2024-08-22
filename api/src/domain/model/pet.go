package model

type Pet struct {
	ID        int64
	Name      string
	PhotoUrls []string
	Status    PetStatus
}

type PetStatus string

const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)
