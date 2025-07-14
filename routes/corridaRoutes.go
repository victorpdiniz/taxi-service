package routes

import (
	"your-app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterCorridaRoutes(app fiber.Router) {
	corridaController := controllers.NewCorridaController()
	app.Post("/corrida/monitorar", corridaController.MonitorarCorrida)
	app.Post("/corrida/finalizar", corridaController.FinalizarCorrida)
	app.Post("/corrida/cancelar-por-excesso-tempo", corridaController.CancelarPorExcessoTempo)
}
