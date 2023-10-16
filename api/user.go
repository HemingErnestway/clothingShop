package api

import (
	"clothingShop/dto"
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func GetIdFromContext(ctx *engine.Context) uint32 {
	path := strings.Split(ctx.Request.URL.Path, "/")
	idString := path[len(path)-1]
	idUint32, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		ctx.Error(400, err.Error())
	}
	return uint32(idUint32)
}

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

	ctx.Print(storage.UserCreate(user))
}

func (h *Handler) UserRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.UserRead(id))
}

func (h *Handler) UsersRead(ctx *engine.Context) {
	ctx.Print(storage.UsersRead())
}

func (h *Handler) UserUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields dto.User
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.UserUpdate(newFields, id))
}

func (h *Handler) UserDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.UserDelete(id))
}
