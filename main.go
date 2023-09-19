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

	shoesCategory := entities.Category{
		Uuid: uuid.New(),
		Name: "Обувь",
	}

	blackColor := entities.Color{
		Uuid: uuid.New(),
		Name: "black",
	}

	summerSeason := entities.Season{
		Uuid: uuid.New(),
		Name: "Лето",
	}

	fourtyFourSize := entities.Size{
		Uuid: uuid.New(),
		Name: "44",
	}

	russiaManufacturer := entities.Manufacturer{
		Uuid: uuid.New(),
		Name: "Россия",
	}

	adidasBrand := entities.Brand{
		Uuid: uuid.New(),
		Name: "Adidas",
	}

	adultAgeGroup := entities.AgeGroup{
		Uuid: uuid.New(),
		Name: "Adult",
	}

	product := entities.Product{
		Uuid:           uuid.New(),
		CategoryId:     shoesCategory.Uuid,
		ColorId:        blackColor.Uuid,
		SeasonId:       summerSeason.Uuid,
		SizeId:         fourtyFourSize.Uuid,
		ManufacturerId: russiaManufacturer.Uuid,
		BrandId:        adidasBrand.Uuid,
		Gender:         entities.Male,
		AgeGroupId:     adultAgeGroup.Uuid,
		PriceRoubles:   3299.90,
	}
	productJSON, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Product %s\n\n", string(productJSON))
	fmt.Println()
}
