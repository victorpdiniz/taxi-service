package routes

import (
	"your-app/controllers"
	"your-app/services"

	"github.com/gofiber/fiber/v2"
)

// SetupCorridaRoutes configura as rotas relacionadas a corridas.
func SetupCorridaRoutes(api fiber.Router, corridaService *services.CorridaService) {
	corridaController := controllers.NewCorridaController(corridaService)

	corridaGroup := api.Group("/corrida")
	corridaGroup.Post("/", corridaController.CriarCorrida)
	corridaGroup.Post("/monitorar", corridaController.MonitorarCorrida)
	corridaGroup.Options("/monitorar", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
}