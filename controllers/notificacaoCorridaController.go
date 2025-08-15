package controllers

import (
	"strconv"
	"strings"
	"taxi-service/models"
	"taxi-service/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ListNotificacoesCorrida - Lista todas as notificações
func ListNotificacoesCorrida(c *fiber.Ctx) error {
	notificacoes, err := services.ListNotificacoesCorrida()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch notificacoes",
		})
	}
	return c.Status(fiber.StatusOK).JSON(notificacoes)
}

// GetNotificacaoCorrida - Busca notificação por ID
func GetNotificacaoCorrida(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	notificacaoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	notificacao, err := services.GetNotificacaoCorrida(uint(notificacaoID))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Notificacao not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch notificacao",
		})
	}

	return c.Status(fiber.StatusOK).JSON(notificacao)
}

// CreateNotificacaoCorrida - Cria nova notificação para motorista
func CreateNotificacaoCorrida(c *fiber.Ctx) error {
	notificacao := new(models.NotificacaoCorrida)

	// Parse do body
	if err := c.BodyParser(notificacao); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Validar MotoristaID
	if notificacao.MotoristaID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":          "MotoristaID is required",
			"received_value": notificacao.MotoristaID,
		})
	}

	// Validar CorridaID
	if notificacao.CorridaID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":          "CorridaID is required",
			"received_value": notificacao.CorridaID,
		})
	}

	// Validar PassageiroNome
	if notificacao.PassageiroNome == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":          "PassageiroNome is required",
			"received_value": notificacao.PassageiroNome,
		})
	}

	// Validar Valor
	if notificacao.Valor <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":          "Valor must be greater than 0",
			"received_value": notificacao.Valor,
		})
	}

	err := services.CreateNotificacaoCorrida(notificacao)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":         "Failed to create notificacao",
			"service_error": err.Error(),
		})
	}


	return c.Status(fiber.StatusCreated).JSON(notificacao)
}

// GetNotificacoesPendentesParaMotorista - Busca notificações pendentes para um motorista específico
func GetNotificacoesPendentesParaMotorista(c *fiber.Ctx) error {
	motoristaIDParam := c.Params("motoristaID")

	if motoristaIDParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "MotoristaID is required",
		})
	}

	motoristaID, err := strconv.ParseUint(motoristaIDParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid MotoristaID format",
		})
	}

	notificacoes, err := services.GetNotificacoesPendentesParaMotorista(uint(motoristaID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch pending notificacoes",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"motorista_id":         motoristaID,
		"pending_count":        len(notificacoes),
		"pending_notificacoes": notificacoes,
	})
}

// AceitarNotificacaoCorrida - Aceita uma notificação de corrida
func AceitarNotificacaoCorrida(c *fiber.Ctx) error {
	notificacaoID := c.Params("id")
	motoristaIDParam := c.Params("motoristaID")

	if notificacaoID == "" || motoristaIDParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "NotificacaoID and MotoristaID are required",
		})
	}

	nID, err := strconv.ParseUint(notificacaoID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid NotificacaoID format",
		})
	}

	mID, err := strconv.ParseUint(motoristaIDParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid MotoristaID format",
		})
	}

	err = services.AceitarNotificacaoCorrida(uint(nID), uint(mID))
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return c.Status(fiber.StatusGone).JSON(fiber.Map{
				"error": "Notificacao expired",
			})
		}
		if strings.Contains(err.Error(), "already processed") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Notificacao already processed",
			})
		}
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Notificacao not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to accept notificacao",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "Notificacao accepted successfully",
		"notificacao_id": nID,
		"motorista_id":   mID,
	})
}

// RecusarNotificacaoCorrida - Recusa uma notificação de corrida
func RecusarNotificacaoCorrida(c *fiber.Ctx) error {
	notificacaoID := c.Params("id")
	motoristaIDParam := c.Params("motoristaID")

	if notificacaoID == "" || motoristaIDParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "NotificacaoID and MotoristaID are required",
		})
	}

	nID, err := strconv.ParseUint(notificacaoID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid NotificacaoID format",
		})
	}

	mID, err := strconv.ParseUint(motoristaIDParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid MotoristaID format",
		})
	}

	err = services.RecusarNotificacaoCorrida(uint(nID), uint(mID))
	if err != nil {
		if strings.Contains(err.Error(), "already processed") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Notificacao already processed",
			})
		}
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Notificacao not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to refuse notificacao",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "Notificacao refused successfully",
		"notificacao_id": nID,
		"motorista_id":   mID,
	})
}

// ExpirarNotificacoesVencidas - Marca como expiradas as notificações que passaram do tempo limite
func ExpirarNotificacoesVencidas(c *fiber.Ctx) error {
	err := services.ExpirarNotificacoesVencidas()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to expire notificacoes",
		})
	}

	return c.JSON(fiber.Map{
		"message":      "Expired notificacoes processed successfully",
		"processed_at": time.Now(),
	})
}
