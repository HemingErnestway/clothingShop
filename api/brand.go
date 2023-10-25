package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) BrandCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var brand entity.Brand
	if err := decoder.Decode(&brand); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.BrandCreate(brand))
}

func (h *Handler) BrandRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.BrandRead(id))
}

func (h *Handler) BrandsRead(ctx *engine.Context) {
	ctx.Print(storage.BrandsRead())
}

func (h *Handler) BrandUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.Brand
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.BrandUpdate(newFields, id))
}

func (h *Handler) BrandDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.BrandDelete(id))
}
