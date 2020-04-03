package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sea_fight/system"
)

type postMatrix struct {
	Range int `json:"range" binding:"required"`
}

func CreateMatrix(app system.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json postMatrix
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		app.CreateMatrix(json.Range)
		c.Status(200)
		//c.JSON(200, gin.H{
		//	"matrix": app.FightMatrix,
		//})
	}
}
