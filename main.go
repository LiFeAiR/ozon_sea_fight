package main

import (
	"github.com/LiFeAiR/ozon_sea_fight/handlers"
	"github.com/LiFeAiR/ozon_sea_fight/system"
	"github.com/gin-gonic/gin"
)

func main() {
	app := system.NewApplication()
	r := gin.Default()
	r.POST("/create-matrix", handlers.CreateMatrix(app))
	r.POST("/ship", handlers.CreateShips(app))
	r.POST("/shot", handlers.MakeShot(app))
	r.GET("/state", handlers.State(app))
	r.POST("/clear", handlers.Clear(app))
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
