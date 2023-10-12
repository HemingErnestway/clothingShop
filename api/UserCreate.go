package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

//type UserCreateRequestBody struct {
//	name     string
//	surname  string
//	email    string
//	login    string
//	password string
//}
//
//func (h *Handler) UserCreate(ctx engine.Context) {
//	var requestBody UserCreateRequestBody
//	if err := c.BindJSON(&requestBody); err != nil {
//		// DO SOMETHING WITH THE ERROR
//	}
//	birthDate, _ := time.Parse("02.01.2006", c.Param("bdate"))
//	usr := entity.User{
//		Name:        requestBody.name,
//		Surname:     requestBody.surname,
//		Email:       requestBody.email,
//		Address:     "",
//		BonusPoints: 0,
//		BirthDate:   birthDate,
//		Login:       requestBody.login,
//		Password:    requestBody.password,
//		Access:      0,
//	}
//	storage.UserNew(usr)
//}

func (h *Handler) UserCreate(ctx engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var user entity.User
	if err := decoder.Decode(&user); err != nil {
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.UserCreate(user))
}

func (h *Handler) UserRead(ctx *engine.Context) {
	path := strings.Split(ctx.Request.URL.Path, "/")
	id := path[len(path)-1]
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	ctx.Print(storage.UserGet(uint32(uid)))
}
