package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sea_fight/system"
)

type postShip struct {
	Coordinates string `binding:"required"`
}

func CreateShips(app system.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		if app.ShipsCreated() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ship repeated"})
			return
		}
		var json postShip
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := app.CreateShips(json.Coordinates)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Status(200)
		//c.JSON(200, gin.H{
		//	"ships": app.Ships,
		//})
	}
}
