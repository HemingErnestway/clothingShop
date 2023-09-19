package entities

type Gender string

const (
	Male   Gender = "male"
	Female        = "female"
)

type Product struct {
	Uuid         UUID
	Category     string
	Color        string
	Season       string
	Size         string
	Manufacturer string
	Brand        string
	Gender       Gender
	Age          string
	PriceRoubles float64
}
