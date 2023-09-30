package api

import (
	"clothingShop/entity"
	"clothingShop/storage"
	"github.com/gin-gonic/gin"
	"time"
)

func UserCreate(c *gin.Context) {
	birthDate, _ := time.Parse("02.01.2006", c.Param("bdate"))
	usr := entity.User{
		Name:        c.Param("name"),
		Surname:     c.Param("surname"),
		Email:       c.Param("email"),
		Address:     "",
		BonusPoints: 0,
		BirthDate:   birthDate,
		Login:       c.Param("login"),
		Password:    c.Param("password"),
		Access:      0,
	}
	storage.UserNew(usr)
}
