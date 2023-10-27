package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) ColorCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var color entity.Color
	if err := decoder.Decode(&color); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.ColorCreate(color))
}

func (h *Handler) ColorRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.ColorRead(id))
}

func (h *Handler) ColorsRead(ctx *engine.Context) {
	ctx.Print(storage.ColorsRead())
}

func (h *Handler) ColorUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.Color
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.ColorUpdate(newFields, id))
}

func (h *Handler) ColorDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.ColorDelete(id))
}
