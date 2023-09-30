package main

import (
	"clothingShop/api"
	"github.com/gin-gonic/gin"
)

func main() {
	//bdate, _ := time.Parse("02.01.2006", "01.12.1992")
	//
	//user := entity.User{
	//	Uuid:        uuid.New(),
	//	Name:        "Иван",
	//	Surname:     "Иванов",
	//	Email:       "ivanovivan@gmail.com",
	//	Address:     "Россия, Московская обл, Долгопрудный, ул Пушкина, 143, 53",
	//	BonusPoints: 200,
	//	BirthDate:   bdate,
	//	Login:       "ivanroflan",
	//	Password:    "qwertyuiop123",
	//	Access:      1,
	//}
	//
	//userJSON, err := json.MarshalIndent(user, "", "  ")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}

	//userMap := map[string]string{}
	//userMap[user.Login] = string(userJSON)

	r := gin.Default()

	r.POST("/UserCreate",
		api.UserCreate)

	r.GET("/readUser/:login", api.UserRead)
	//var v interface{}
	//json.Unmarshal(userJSON, &v)
	//data := v.(map[string]interface{})

	//usersJSON, err := json.MarshalIndent()
	//c.Data(http.StatusOK, gin.MIMEJSON)

	r.Run()

	//
	//shoesCategory := entity.Category{
	//	Uuid: uuid.New(),
	//	Name: "Обувь",
	//}
	//
	//blackColor := entity.Color{
	//	Uuid: uuid.New(),
	//	Name: "black",
	//}
	//
	//summerSeason := entity.Season{
	//	Uuid: uuid.New(),
	//	Name: "Лето",
	//}
	//
	//fourtyFourSize := entity.Size{
	//	Uuid: uuid.New(),
	//	Name: "44",
	//}
	//
	//russiaManufacturer := entity.Manufacturer{
	//	Uuid: uuid.New(),
	//	Name: "Россия",
	//}
	//
	//adidasBrand := entity.Brand{
	//	Uuid: uuid.New(),
	//	Name: "Adidas",
	//}
	//
	//adultAgeGroup := entity.AgeGroup{
	//	Uuid: uuid.New(),
	//	Name: "Adult",
	//}
	//
	//product := entity.Product{
	//	Uuid:           uuid.New(),
	//	CategoryId:     shoesCategory.Uuid,
	//	ColorId:        blackColor.Uuid,
	//	SeasonId:       summerSeason.Uuid,
	//	SizeId:         fourtyFourSize.Uuid,
	//	ManufacturerId: russiaManufacturer.Uuid,
	//	BrandId:        adidasBrand.Uuid,
	//	Gender:         entity.Male,
	//	AgeGroupId:     adultAgeGroup.Uuid,
	//	PriceRoubles:   3299.90,
	//}
	//productJSON, err := json.MarshalIndent(product, "", "  ")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//fmt.Printf("Product %s\n\n", string(productJSON))
}
