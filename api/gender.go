package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) GenderCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var gender entity.Gender
	if err := decoder.Decode(&gender); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.GenderCreate(gender))
}

func (h *Handler) GenderRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.GenderRead(id))
}

func (h *Handler) GendersRead(ctx *engine.Context) {
	ctx.Print(storage.GendersRead())
}

func (h *Handler) GenderUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.Gender
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.GenderUpdate(newFields, id))
}

func (h *Handler) GenderDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.GenderDelete(id))
}
