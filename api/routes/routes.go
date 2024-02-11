package routes

import (
	"github.com/Ellathet/carnaval/api/ping"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", ping.Ping)

	return r
}
