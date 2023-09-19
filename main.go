package main

import (
	"awesomeProject/entities"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"time"

	"fmt"
)

func main() {
	fmt.Println("Смирнов Антон Леонидович")
	fmt.Println("Интернет-магазин одежды\n")
	bdate, _ := time.Parse("02.01.2006", "01.12.1992")
	customer := entities.Customer{
		Uuid:        uuid.New(),
		Name:        "Иван",
		Surname:     "Иванов",
		Email:       "ivanovivan@gmail.com",
		Address:     "Россия, Московская обл, Долгопрудный, ул Пушкина, 143, 53",
		BonusPoints: 200,
		BirthDate:   bdate,
	}
	customerJSON, err := json.MarshalIndent(customer, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Customer %s\n\n", string(customerJSON))

	administrator := entities.Administrator{
		Uuid:    entities.UUID{2222, 3333, 1111, 4444},
		Name:    "Петр",
		Surname: "Петрович",
		Email:   "petrovpetr@gmail.com",
		Address: "Россия, Московская обл, Долгопрудный, ул Колотушкина, 53, 123",
	}
	administratorJSON, err := json.MarshalIndent(administrator, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Administrator %s\n\n", string(administratorJSON))

	product := entities.Product{
		Uuid:         entities.UUID{1234, 1234, 1234, 1234},
		Category:     "Обувь",
		Color:        "Черный",
		Season:       "Лето",
		Size:         "44",
		Manufacturer: "Россия",
		Brand:        "Adidas",
		Gender:       entities.Male,
		Age:          "Adult",
		PriceRoubles: 3299.90,
	}
	productJSON, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Administrator %s\n\n", string(productJSON))
}
