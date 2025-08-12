package controllers

import (
	"fmt" // Adicionar este import
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
	fmt.Printf("\n🎯 [CONTROLLER] CreateNotificacaoCorrida - INÍCIO\n")

	// Debug: Headers da requisição
	fmt.Printf("📡 [DEBUG] Headers da requisição:\n")
	fmt.Printf("   - Content-Type: %s\n", c.Get("Content-Type"))
	fmt.Printf("   - User-Agent: %s\n", c.Get("User-Agent"))
	fmt.Printf("   - Content-Length: %s\n", c.Get("Content-Length"))

	// Debug: Método e URL
	fmt.Printf("🔗 [DEBUG] Requisição:\n")
	fmt.Printf("   - Método: %s\n", c.Method())
	fmt.Printf("   - URL: %s\n", c.OriginalURL())
	fmt.Printf("   - IP do cliente: %s\n", c.IP())

	// Debug: Body raw antes do parse
	bodyBytes := c.Body()
	fmt.Printf("📄 [DEBUG] Body raw recebido: %s\n", string(bodyBytes))
	fmt.Printf("📏 [DEBUG] Tamanho do body: %d bytes\n", len(bodyBytes))

	// Criar nova instância da notificação
	notificacao := new(models.NotificacaoCorrida)
	fmt.Printf("🆕 [DEBUG] Nova instância de NotificacaoCorrida criada\n")

	// Parse do body
	fmt.Printf("🔄 [DEBUG] Iniciando parse do body...\n")
	if err := c.BodyParser(notificacao); err != nil {
		fmt.Printf("❌ [ERROR] Erro no parse do body: %v\n", err)
		fmt.Printf("📋 [ERROR] Tipo do erro: %T\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}
	fmt.Printf("✅ [DEBUG] Parse do body realizado com sucesso\n")

	// Debug: Dados parseados
	fmt.Printf("📊 [DEBUG] Dados parseados da requisição:\n")
	fmt.Printf("   - CorridaID: %d\n", notificacao.CorridaID)
	fmt.Printf("   - MotoristaID: %d\n", notificacao.MotoristaID)
	fmt.Printf("   - PassageiroNome: '%s'\n", notificacao.PassageiroNome)
	fmt.Printf("   - Origem: '%s'\n", notificacao.Origem)
	fmt.Printf("   - Destino: '%s'\n", notificacao.Destino)
	fmt.Printf("   - Valor: %.2f\n", notificacao.Valor)
	fmt.Printf("   - DistanciaKm: %.2f\n", notificacao.DistanciaKm)
	fmt.Printf("   - TempoEstimado: %s\n", notificacao.TempoEstimado)

	// Validações com debug detalhado
	fmt.Printf("🔍 [DEBUG] Iniciando validações...\n")

	// Validar MotoristaID
	fmt.Printf("🔎 [VALIDATION] Verificando MotoristaID...\n")
	if notificacao.MotoristaID == 0 {
		fmt.Printf("❌ [VALIDATION ERROR] MotoristaID é obrigatório (recebido: %d)\n", notificacao.MotoristaID)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":          "MotoristaID is required",
			"received_value": notificacao.MotoristaID,
		})
	}
	fmt.Printf("✅ [VALIDATION] MotoristaID válido: %d\n", notificacao.MotoristaID)

	// Validar CorridaID
	fmt.Printf("🔎 [VALIDATION] Verificando CorridaID...\n")
	if notificacao.CorridaID == 0 {
		fmt.Printf("❌ [VALIDATION ERROR] CorridaID é obrigatório (recebido: %d)\n", notificacao.CorridaID)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":          "CorridaID is required",
			"received_value": notificacao.CorridaID,
		})
	}
	fmt.Printf("✅ [VALIDATION] CorridaID válido: %d\n", notificacao.CorridaID)

	// Validar PassageiroNome
	fmt.Printf("🔎 [VALIDATION] Verificando PassageiroNome...\n")
	if notificacao.PassageiroNome == "" {
		fmt.Printf("❌ [VALIDATION ERROR] PassageiroNome é obrigatório (recebido: '%s')\n", notificacao.PassageiroNome)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":          "PassageiroNome is required",
			"received_value": notificacao.PassageiroNome,
		})
	}
	fmt.Printf("✅ [VALIDATION] PassageiroNome válido: '%s'\n", notificacao.PassageiroNome)

	// Validar Valor
	fmt.Printf("🔎 [VALIDATION] Verificando Valor...\n")
	if notificacao.Valor <= 0 {
		fmt.Printf("❌ [VALIDATION ERROR] Valor deve ser maior que 0 (recebido: %.2f)\n", notificacao.Valor)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":          "Valor must be greater than 0",
			"received_value": notificacao.Valor,
		})
	}
	fmt.Printf("✅ [VALIDATION] Valor válido: %.2f\n", notificacao.Valor)

	fmt.Printf("🎉 [DEBUG] Todas as validações passaram com sucesso!\n")

	// Chamar o service
	fmt.Printf("📞 [DEBUG] Chamando services.CreateNotificacaoCorrida...\n")
	startTime := time.Now()

	err := services.CreateNotificacaoCorrida(notificacao)

	duration := time.Since(startTime)
	fmt.Printf("⏱️  [DEBUG] Tempo de execução do service: %v\n", duration)

	if err != nil {
		fmt.Printf("❌ [SERVICE ERROR] Erro no service: %v\n", err)
		fmt.Printf("📋 [SERVICE ERROR] Tipo do erro: %T\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":         "Failed to create notificacao",
			"service_error": err.Error(),
		})
	}
	fmt.Printf("✅ [DEBUG] Service executado com sucesso!\n")

	// Debug: Dados finais que serão retornados
	fmt.Printf("📤 [DEBUG] Dados que serão retornados:\n")
	fmt.Printf("   - ID: %d\n", notificacao.ID)
	fmt.Printf("   - CorridaID: %d\n", notificacao.CorridaID)
	fmt.Printf("   - MotoristaID: %d\n", notificacao.MotoristaID)
	fmt.Printf("   - Status: %s\n", notificacao.Status)
	fmt.Printf("   - CreatedAt: %s\n", notificacao.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("   - ExpiraEm: %s\n", notificacao.ExpiraEm.Format("2006-01-02 15:04:05"))

	fmt.Printf("🎯 [CONTROLLER] CreateNotificacaoCorrida - SUCESSO!\n")
	fmt.Printf("📋 [RESPONSE] Status: 201 Created\n")
	fmt.Printf("🔚 [CONTROLLER] CreateNotificacaoCorrida - FIM\n\n")

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
