package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CountryCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var country entity.Country
	if err := decoder.Decode(&country); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.CountryCreate(country))
}

func (h *Handler) CountryRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.CountryRead(id))
}

func (h *Handler) CountriesRead(ctx *engine.Context) {
	ctx.Print(storage.CountriesRead())
}

func (h *Handler) CountryUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.Country
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.CountryUpdate(newFields, id))
}

func (h *Handler) CountryDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.CountryDelete(id))
}
