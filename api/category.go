package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CategoryCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var category entity.Category
	if err := decoder.Decode(&category); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.CategoryCreate(category))
}

func (h *Handler) CategoryRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.CategoryRead(id))
}

func (h *Handler) CategoriesRead(ctx *engine.Context) {
	ctx.Print(storage.CategoriesRead())
}

func (h *Handler) CategoryUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.Category
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.CategoryUpdate(newFields, id))
}

func (h *Handler) CategoryDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.CategoryDelete(id))
}
