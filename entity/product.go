package entity

import "github.com/google/uuid"

type Gender string

const (
	Male   Gender = "male"
	Female        = "female"
)

type Product struct {
	Uuid           uuid.UUID
	CategoryId     uuid.UUID
	ColorId        uuid.UUID
	SeasonId       uuid.UUID
	SizeId         uuid.UUID
	ManufacturerId uuid.UUID
	BrandId        uuid.UUID
	Gender         Gender
	AgeGroupId     uuid.UUID
	PriceRoubles   float64
}
