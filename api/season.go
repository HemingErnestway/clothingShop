package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) SeasonCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var season entity.Season
	if err := decoder.Decode(&season); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.SeasonCreate(season))
}

func (h *Handler) SeasonRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.SeasonRead(id))
}

func (h *Handler) SeasonsRead(ctx *engine.Context) {
	ctx.Print(storage.SeasonsRead())
}

func (h *Handler) SeasonUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.Season
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.SeasonUpdate(newFields, id))
}

func (h *Handler) SeasonDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.SeasonDelete(id))
}
