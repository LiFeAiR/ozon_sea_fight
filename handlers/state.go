package handlers

import (
	"github.com/gin-gonic/gin"
	"sea_fight/system"
)

func State(app system.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, app.State())
	}
}
