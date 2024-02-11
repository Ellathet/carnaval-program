package main

import (
	"github.com/Ellathet/carnaval/api/routes"
	config "github.com/Ellathet/carnaval/configs"
	"github.com/Ellathet/carnaval/models"
	"github.com/Ellathet/carnaval/utils"
)

func main() {
	r := routes.SetupRouter()

	utils.DB.AutoMigrate(&models.Block{})

	r.Run(":" + config.Env.Port)
}
