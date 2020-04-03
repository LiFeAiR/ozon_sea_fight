package handlers

import (
	"github.com/gin-gonic/gin"
	"sea_fight/system"
)

func Clear(app system.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		app.Clear()
		c.JSON(200, app.State())
	}
}
