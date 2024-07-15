package handlers

import (
	"github.com/LiFeAiR/ozon_sea_fight/system"
	"github.com/gin-gonic/gin"
)

func Clear(app system.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		app.Clear()
		c.JSON(200, app.State())
	}
}
