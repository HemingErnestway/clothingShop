package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) UserCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var user entity.User
	if err := decoder.Decode(&user); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	user.Access = 0
	user.BonusPoints = 0
	userS := storage.UserCreate(user)
	fmt.Println("UserCreate", user, userS)
	ctx.Print(userS)
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
