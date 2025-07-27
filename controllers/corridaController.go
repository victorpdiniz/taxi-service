package controllers

import (
	"taxi_service/models"
	"taxi_service/services"

	"github.com/gofiber/fiber/v2"
	"strconv"
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

func (cc *CorridaController) AvaliarCorrida(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	var input struct {
		Nota int `json:"nota"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON inválido"})
	}

	if err := services.AvaliarCorrida(id, input.Nota); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "avaliado"})
}

func (cc *CorridaController) CriarCorrida(c *fiber.Ctx) error {
	var corrida models.Corrida

	if err := c.BodyParser(&corrida); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "JSON inválido",
		})
	}

	// Salva na slice em memória
	cc.service.AdicionarCorrida(corrida)

	return c.Status(fiber.StatusCreated).JSON(corrida)
}

func (cc *CorridaController) ListarCorridas(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(services.GetCorridas())
}
