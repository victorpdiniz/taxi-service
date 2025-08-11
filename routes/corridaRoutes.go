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
	corridaGroup.Get("/:id", corridaController.GetCorrida) // Nova rota
	corridaGroup.Post("/monitorar", corridaController.MonitorarCorrida)
	corridaGroup.Put("/:id/aceitar", corridaController.AceitarCorrida)
	corridaGroup.Put("/:id/posicao", corridaController.AtualizarPosicao)
	corridaGroup.Post("/:id/cancelar", corridaController.CancelarCorrida) // Nova rota
	corridaGroup.Post("/:id/finalizar", corridaController.FinalizarCorrida) // Nova rota

	// Manter a rota OPTIONS para o CORS
	corridaGroup.Options("/monitorar", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
}
