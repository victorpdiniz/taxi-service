package services

import (
	"time"
	"your-app/models"

	"github.com/gofiber/fiber/v2"
)

// NotificacaoService é responsável por enviar notificações
func NotificarPassageiro(ctx *fiber.Ctx, passageiroID int, mensagem string) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"notificacao": mensagem, "passageiro_id": passageiroID})
}

func NotificarMotorista(ctx *fiber.Ctx, motoristaID int, mensagem string) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"notificacao": mensagem, "motorista_id": motoristaID})
}

// BonusService aplica bônus ao motorista
func AplicarBonus(motoristaID int) {
	// Implementação fictícia: aplicar bônus
	// Exemplo: atualizar saldo do motorista
}

// CorridaService contém regras de negócio para corrida
type CorridaService struct{}

func (s *CorridaService) VerificarTempoCorrida(ctx *fiber.Ctx, corrida *models.Corrida) {
	diferenca := corrida.TempoDecorrido - corrida.TempoEstimado

	if corrida.Status == models.StatusEmAndamento {
		if diferenca > 0 && diferenca <= 15 {
			corrida.Status = models.StatusAtrasado
			NotificarPassageiro(ctx, corrida.PassageiroID, "O motorista está atrasado.")
			NotificarMotorista(ctx, corrida.MotoristaID, "Tag: atrasado")
		} else if diferenca > 15 {
			corrida.Status = models.StatusCanceladaPorExcessoTempo
			NotificarMotorista(ctx, corrida.MotoristaID, "Corrida cancelada por excesso de tempo")
			// Corrida cancelada automaticamente
		}
	}
}

func (s *CorridaService) FinalizarCorrida(ctx *fiber.Ctx, corrida *models.Corrida) {
	diferenca := corrida.TempoDecorrido - corrida.TempoEstimado
	now := time.Now()
	corrida.DataFim = &now

	if diferenca < 0 {
		corrida.Status = models.StatusConcluidaAntecedencia
		NotificarMotorista(ctx, corrida.MotoristaID, "Parabéns! Corrida concluída com antecedência.")
		AplicarBonus(corrida.MotoristaID)
		corrida.BonusAplicado = true
	} else if diferenca == 0 || diferenca > 0 && diferenca <= 15 {
		corrida.Status = models.StatusConcluidaNoTempo
		// Nenhuma penalização ou bônus
	}
}
