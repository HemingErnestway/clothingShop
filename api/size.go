package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) SizeCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var size entity.Size
	if err := decoder.Decode(&size); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.SizeCreate(size))
}

func (h *Handler) SizeRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.SizeRead(id))
}

func (h *Handler) SizesRead(ctx *engine.Context) {
	ctx.Print(storage.SizesRead())
}

func (h *Handler) SizeUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.Size
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.SizeUpdate(newFields, id))
}

func (h *Handler) SizeDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.SizeDelete(id))
}
