package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) AgeGroupCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var ageGroup entity.AgeGroup
	if err := decoder.Decode(&ageGroup); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.AgeGroupCreate(ageGroup))
}

func (h *Handler) AgeGroupRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.AgeGroupRead(id))
}

func (h *Handler) AgeGroupsRead(ctx *engine.Context) {
	ctx.Print(storage.AgeGroupsRead())
}

func (h *Handler) AgeGroupUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.AgeGroup
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.AgeGroupUpdate(newFields, id))
}

func (h *Handler) AgeGroupDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.AgeGroupDelete(id))
}
