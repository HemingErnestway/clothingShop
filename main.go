package main

import (
	"clothingShop/entities"
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

	user := entities.User{
		Uuid:        uuid.New(),
		Name:        "Иван",
		Surname:     "Иванов",
		Email:       "ivanovivan@gmail.com",
		Address:     "Россия, Московская обл, Долгопрудный, ул Пушкина, 143, 53",
		BonusPoints: 200,
		BirthDate:   bdate,
		Login:       "ivanroflan",
		Password:    "qwertyuiop123",
		Access:      1,
	}
	userJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("User %s\n\n", string(userJSON))

	product := entities.Product{
		Uuid:         uuid.New(),
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
	fmt.Printf("Product %s\n\n", string(productJSON))
}
