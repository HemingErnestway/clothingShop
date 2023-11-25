package api

import (
	"clothingShop/db"
	"clothingShop/dto"
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"
	"time"
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

func (h *Handler) UserAuth(ctx *engine.Context) {
	userDto, err := engine.ToStruct[dto.UserAuth](ctx)
	if err != nil {
		ctx.Error(400, "Bad user data")
	}

	var user entity.User
	db.DB().Table(user.TableName()).Where("login = ? and password = ?",
		userDto.Login, userDto.Password).Find(&user)

	if user.Uuid == 0 {
		ctx.Error(401, "Bad auth")
		return
	}

	token := entity.Token{
		Uid:     user.Uuid,
		Token:   uuid.NewString(),
		Expired: time.Now().Add(1 * time.Hour),
	}

	db.DB().Save(&token)
	ctx.Print(token)
}
