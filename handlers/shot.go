package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sea_fight/system"
)

type postShot struct {
	Coordinates string `json:"coord" binding:"required"`
}

func MakeShot(app system.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json postShot
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := app.MakeShot(json.Coordinates)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, response)
	}
}
