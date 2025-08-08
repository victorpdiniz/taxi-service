package controllers

import (
	"taxi-service/models"
	"taxi-service/services"

	"github.com/gofiber/fiber/v2"
)

type CorridaController struct {
	service *services.CorridaService
}

func NewCorridaController() *CorridaController {
	return &CorridaController{service: &services.CorridaService{}}
}

// POST /corrida/monitorar
func (cc *CorridaController) MonitorarCorrida(c *fiber.Ctx) error {
	var corrida models.Corrida
	if err := c.BodyParser(&corrida); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	cc.service.VerificarTempoCorrida(c, &corrida)
	return c.Status(fiber.StatusOK).JSON(corrida)
}

// POST /corrida/finalizar
func (cc *CorridaController) FinalizarCorrida(c *fiber.Ctx) error {
	var corrida models.Corrida
	if err := c.BodyParser(&corrida); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	cc.service.FinalizarCorrida(c, &corrida)
	return c.Status(fiber.StatusOK).JSON(corrida)
}

// POST /corrida/cancelar-por-excesso-tempo
func (cc *CorridaController) CancelarPorExcessoTempo(c *fiber.Ctx) error {
	var corrida models.Corrida
	if err := c.BodyParser(&corrida); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if corrida.TempoDecorrido-corrida.TempoEstimado > 15 {
		corrida.Status = models.StatusCanceladaPorExcessoTempo
		services.NotificarMotorista(c, corrida.MotoristaID, "Corrida cancelada por excesso de tempo")
	}
	return c.Status(fiber.StatusOK).JSON(corrida)
}
