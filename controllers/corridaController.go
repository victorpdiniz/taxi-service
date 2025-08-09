package controllers

import (
	"your-app/models"
	"your-app/services"

	"github.com/gofiber/fiber/v2"
)

// CorridaController gerencia as requisições HTTP para corridas.
type CorridaController struct {
	service *services.CorridaService
}

// NewCorridaController cria uma nova instância de CorridaController.
func NewCorridaController(service *services.CorridaService) *CorridaController {
	return &CorridaController{service: service}
}

// CriarCorrida (POST /corrida) cria uma nova corrida.
func (cc *CorridaController) CriarCorrida(c *fiber.Ctx) error {
	var corridaInput models.Corrida
	if err := c.BodyParser(&corridaInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Não foi possível decodificar o corpo da requisição"})
	}

	// Validação básica
	if corridaInput.PassageiroID == 0 || corridaInput.TempoEstimado <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "PassageiroID e TempoEstimado são obrigatórios"})
	}

	corrida, err := cc.service.CriarNovaCorrida(corridaInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(corrida)
}

// GetCorrida (GET /corrida/:id) busca o status de uma corrida.
// MonitorarCorrida (POST /corrida/monitorar) monitora uma corrida.
func (cc *CorridaController) MonitorarCorrida(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}