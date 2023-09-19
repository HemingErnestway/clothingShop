package entities

import "github.com/google/uuid"

type Category struct {
	Uuid uuid.UUID
	Name string
}
