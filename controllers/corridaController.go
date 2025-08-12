package controllers

import (
	"log"
	"strconv"
	"taxi-service/models"
	"taxi-service/services"

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
	log.Printf("[DEBUG] CriarCorrida - Headers: %v", c.GetReqHeaders())
	log.Printf("[DEBUG] CriarCorrida - Body: %s", string(c.Body()))

	var corridaInput models.Corrida
	if err := c.BodyParser(&corridaInput); err != nil {
		log.Printf("[ERROR] CriarCorrida - Erro ao fazer parse: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Não foi possível decodificar o corpo da requisição"})
	}

	log.Printf("[DEBUG] CriarCorrida - Dados parseados: %+v", corridaInput)

	if corridaInput.PassageiroID == 0 {
		log.Printf("[ERROR] CriarCorrida - PassageiroID vazio")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "PassageiroID é obrigatório"})
	}

	corrida, err := cc.service.CriarNovaCorrida(corridaInput)
	if err != nil {
		log.Printf("[ERROR] CriarCorrida - Erro no service: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("[DEBUG] CriarCorrida - Corrida criada: %+v", corrida)

	cc.service.AdicionarCorrida(corridaInput)

	log.Printf("[SUCCESS] CriarCorrida - Corrida criada com sucesso, ID: %v", corrida)
	return c.Status(fiber.StatusCreated).JSON(corrida)
}

// GetCorrida (GET /corrida/:id) busca o status de uma corrida.
func (cc *CorridaController) GetCorrida(c *fiber.Ctx) error {
	idParam := c.Params("id")
	log.Printf("[DEBUG] GetCorrida - ID recebido: %s", idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("[ERROR] GetCorrida - ID inválido: %s", idParam)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID da corrida inválido"})
	}

	corrida, err := cc.service.GetCorridaPorID(id)
	if err != nil {
		log.Printf("[ERROR] GetCorrida - Corrida não encontrada: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("[SUCCESS] GetCorrida - Corrida encontrada: %+v", corrida)
	return c.JSON(corrida)
}

// MonitorarCorrida (POST /corrida/monitorar) monitora uma corrida.
func (cc *CorridaController) MonitorarCorrida(c *fiber.Ctx) error {
	log.Printf("[DEBUG] MonitorarCorrida - Chamada recebida")
	return c.SendStatus(fiber.StatusOK)
}

// AceitarCorrida (PUT /corrida/:id/aceitar) permite que um motorista aceite uma corrida.
func (cc *CorridaController) AceitarCorrida(c *fiber.Ctx) error {
	idParam := c.Params("id")
	log.Printf("[DEBUG] AceitarCorrida - ID: %s, Body: %s", idParam, string(c.Body()))

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("[ERROR] AceitarCorrida - ID inválido: %s", idParam)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID da corrida inválido"})
	}

	var body struct {
		MotoristaID int `json:"motoristaId"`
	}

	if err := c.BodyParser(&body); err != nil {
		log.Printf("[ERROR] AceitarCorrida - Erro no parse: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Corpo da requisição inválido"})
	}

	log.Printf("[DEBUG] AceitarCorrida - CorridaID: %d, MotoristaID: %d", id, body.MotoristaID)

	if err := cc.service.AceitarCorrida(id, body.MotoristaID); err != nil {
		log.Printf("[ERROR] AceitarCorrida - Erro no service: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("[SUCCESS] AceitarCorrida - Corrida aceita")
	return c.SendStatus(fiber.StatusOK)
}

// AtualizarPosicao (PUT /corrida/:id/posicao) atualiza a posição do motorista.
func (cc *CorridaController) AtualizarPosicao(c *fiber.Ctx) error {
	idParam := c.Params("id")
	log.Printf("[DEBUG] AtualizarPosicao - ID: %s, Body: %s", idParam, string(c.Body()))

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("[ERROR] AtualizarPosicao - ID inválido: %s", idParam)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID da corrida inválido"})
	}

	var body struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	if err := c.BodyParser(&body); err != nil {
		log.Printf("[ERROR] AtualizarPosicao - Erro no parse: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Corpo da requisição inválido"})
	}

	log.Printf("[DEBUG] AtualizarPosicao - CorridaID: %d, Lat: %f, Lng: %f", id, body.Lat, body.Lng)

	if err := cc.service.AtualizarPosicao(id, body.Lat, body.Lng); err != nil {
		log.Printf("[ERROR] AtualizarPosicao - Erro no service: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("[SUCCESS] AtualizarPosicao - Posição atualizada")
	return c.SendStatus(fiber.StatusOK)
}

// CancelarCorrida (POST /corrida/:id/cancelar) cancela uma corrida.
func (cc *CorridaController) CancelarCorrida(c *fiber.Ctx) error {
	idParam := c.Params("id")
	log.Printf("[DEBUG] CancelarCorrida - ID: %s", idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("[ERROR] CancelarCorrida - ID inválido: %s", idParam)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID da corrida inválido"})
	}

	if err := cc.service.CancelarCorrida(id); err != nil {
		log.Printf("[ERROR] CancelarCorrida - Erro no service: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("[SUCCESS] CancelarCorrida - Corrida cancelada, ID: %d", id)
	return c.SendStatus(fiber.StatusOK)
}

// FinalizarCorrida (POST /corrida/:id/finalizar) finaliza uma corrida.
func (cc *CorridaController) FinalizarCorrida(c *fiber.Ctx) error {
	idParam := c.Params("id")
	log.Printf("[DEBUG] FinalizarCorrida - ID: %s", idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("[ERROR] FinalizarCorrida - ID inválido: %s", idParam)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID da corrida inválido"})
	}

	if err := cc.service.FinalizarCorrida(id); err != nil {
		log.Printf("[ERROR] FinalizarCorrida - Erro no service: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("[SUCCESS] FinalizarCorrida - Corrida finalizada, ID: %d", id)
	return c.SendStatus(fiber.StatusOK)
}

func (cc *CorridaController) AvaliarCorrida(c *fiber.Ctx) error {
	idStr := c.Params("id")
	log.Printf("[DEBUG] AvaliarCorrida - ID: %s, Body: %s", idStr, string(c.Body()))

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[ERROR] AvaliarCorrida - ID inválido: %s", idStr)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
	}

	var input struct {
		Nota int `json:"nota"`
	}
	if err := c.BodyParser(&input); err != nil {
		log.Printf("[ERROR] AvaliarCorrida - Erro no parse: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON inválido"})
	}

	log.Printf("[DEBUG] AvaliarCorrida - CorridaID: %d, Nota: %d", id, input.Nota)

	if err := services.AvaliarCorrida(id, input.Nota); err != nil {
		log.Printf("[ERROR] AvaliarCorrida - Erro no service: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	log.Printf("[SUCCESS] AvaliarCorrida - Corrida avaliada")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "avaliado"})
}

func (cc *CorridaController) ListarCorridas(c *fiber.Ctx) error {
	log.Printf("[DEBUG] ListarCorridas - Chamada recebida")
	corridas := services.GetCorridas()
	log.Printf("[DEBUG] ListarCorridas - Retornando %d corridas", len(corridas))
	return c.Status(fiber.StatusOK).JSON(corridas)
}

func (cc *CorridaController) Service() *services.CorridaService {
	return cc.service
}
