package api

import (
	"clothingShop/storage"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserRead(c *gin.Context) {
	userJSON, err := json.MarshalIndent(
		storage.UserFind(c.Param("login")), "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	c.Data(http.StatusOK, gin.MIMEJSON, userJSON)
}
