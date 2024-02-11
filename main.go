package main

import (
	"github.com/Ellathet/carnaval/api/routes"
	config "github.com/Ellathet/carnaval/configs"
)

func main() {
	r := routes.SetupRouter()

	r.Run(":" + config.Env.Port)
}
