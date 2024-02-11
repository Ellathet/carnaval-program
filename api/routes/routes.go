package routes

import (
	"github.com/Ellathet/carnaval/api/block"
	"github.com/Ellathet/carnaval/api/ping"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", ping.Ping)

	r.POST("/block", block.CreateBlock)
	r.GET("/block/:id", block.GetBlock)
	r.PUT("/block/:id", block.UpdateBlock)
	r.DELETE("/block/:id", block.DeleteBlock)

	return r
}
