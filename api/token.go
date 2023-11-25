package api

import (
	"clothingShop/engine"
	"clothingShop/entity"
	"clothingShop/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) TokenCreate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var token entity.Token
	if err := decoder.Decode(&token); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Print(storage.TokenCreate(token))
}

func (h *Handler) TokenRead(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.TokenRead(id))
}

func (h *Handler) TokensRead(ctx *engine.Context) {
	ctx.Print(storage.TokensRead())
}

func (h *Handler) TokenUpdate(ctx *engine.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var newFields entity.Token
	if err := decoder.Decode(&newFields); err != nil {
		fmt.Println("Error", err)
		ctx.Error(http.StatusBadRequest, err.Error())
		return
	}
	id := GetIdFromContext(ctx)
	ctx.Print(storage.TokenUpdate(newFields, id))
}

func (h *Handler) TokenDelete(ctx *engine.Context) {
	id := GetIdFromContext(ctx)
	ctx.Print(storage.TokenDelete(id))
}
