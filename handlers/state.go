package handlers

import (
	"github.com/LiFeAiR/ozon_sea_fight/system"
	"github.com/gin-gonic/gin"
)

func State(app system.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, app.State())
	}
}
