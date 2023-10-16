package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) ProductCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var product entity.Product
	if err := decoder.Decode(&product); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.ProductCreate(product))
}

func (h *Handler) ProductRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.ProductRead(id))
}

func (h *Handler) ProductsRead(ctx *engine.Context) {
	ctx.Print(storage.ProductsRead())
}

func (h *Handler) ProductUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.Product
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.ProductUpdate(newFields, id))
}

func (h *Handler) ProductDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.ProductDelete(id))
}
