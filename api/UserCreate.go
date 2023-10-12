package api

import (
	"clothingShop/entity"
	"clothingShop/storage"
	"github.com/gin-gonic/gin"
	"time"
)

type UserCreateRequestBody struct {
	name     string
	surname  string
	email    string
	login    string
	password string
}

func UserCreate(c *gin.Context) {
	var requestBody UserCreateRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
	}
	birthDate, _ := time.Parse("02.01.2006", c.Param("bdate"))
	usr := entity.User{
		Name:        requestBody.name,
		Surname:     requestBody.surname,
		Email:       requestBody.email,
		Address:     "",
		BonusPoints: 0,
		BirthDate:   birthDate,
		Login:       requestBody.login,
		Password:    requestBody.password,
		Access:      0,
	}
	storage.UserNew(usr)
}
