package api

import (
	"clothingShop/dto"
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CartCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var cart entity.Cart
	if err := decoder.Decode(&cart); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.CartCreate(cart))
}

func (h *Handler) CartRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.CartRead(id))
}

func (h *Handler) CartUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields dto.Cart
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.CartUpdate(newFields, id))
}

func (h *Handler) CartDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.CartDelete(id))
}
