package main

import (
	"clothingShop/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"encoding/json"
	"log"
	"net/http"
	"time"
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

	userMap := map[string]string{}
	userMap[user.Login] = string(userJSON)

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		//var v interface{}
		//json.Unmarshal(userJSON, &v)
		//data := v.(map[string]interface{})
		c.Data(http.StatusOK, gin.MIMEJSON, userJSON)
	})

	r.GET("/echo/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, name)
	})

	r.Run()

	//
	//shoesCategory := entities.Category{
	//	Uuid: uuid.New(),
	//	Name: "Обувь",
	//}
	//
	//blackColor := entities.Color{
	//	Uuid: uuid.New(),
	//	Name: "black",
	//}
	//
	//summerSeason := entities.Season{
	//	Uuid: uuid.New(),
	//	Name: "Лето",
	//}
	//
	//fourtyFourSize := entities.Size{
	//	Uuid: uuid.New(),
	//	Name: "44",
	//}
	//
	//russiaManufacturer := entities.Manufacturer{
	//	Uuid: uuid.New(),
	//	Name: "Россия",
	//}
	//
	//adidasBrand := entities.Brand{
	//	Uuid: uuid.New(),
	//	Name: "Adidas",
	//}
	//
	//adultAgeGroup := entities.AgeGroup{
	//	Uuid: uuid.New(),
	//	Name: "Adult",
	//}
	//
	//product := entities.Product{
	//	Uuid:           uuid.New(),
	//	CategoryId:     shoesCategory.Uuid,
	//	ColorId:        blackColor.Uuid,
	//	SeasonId:       summerSeason.Uuid,
	//	SizeId:         fourtyFourSize.Uuid,
	//	ManufacturerId: russiaManufacturer.Uuid,
	//	BrandId:        adidasBrand.Uuid,
	//	Gender:         entities.Male,
	//	AgeGroupId:     adultAgeGroup.Uuid,
	//	PriceRoubles:   3299.90,
	//}
	//productJSON, err := json.MarshalIndent(product, "", "  ")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//fmt.Printf("Product %s\n\n", string(productJSON))
}
