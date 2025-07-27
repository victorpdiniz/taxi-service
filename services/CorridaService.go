package services

import (
	"taxi_service/models"
	"time"

	"github.com/gofiber/fiber/v2"

	"errors"

	"encoding/json"
	"os"
	"log"
)

var corridas []models.Corrida

type CorridaService struct {
	Corridas []models.Corrida // Mock de "banco de dados" em memória
}

// NotificacaoService é responsável por enviar notificações
func NotificarPassageiro(ctx *fiber.Ctx, passageiroID int, mensagem string) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"notificacao": mensagem, "passageiro_id": passageiroID})
}

func NotificarMotorista(ctx *fiber.Ctx, motoristaID int, mensagem string) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"notificacao": mensagem, "motorista_id": motoristaID})
}

func AplicarBonus(corrida *models.Corrida) {
	// Aplica 10% de bônus sobre o preço da corrida, se o campo existir
	if corrida != nil {
		corrida.Preco = corrida.Preco + (corrida.Preco * 0.10)
	}

}

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
			// Corrida cancelada: Esperar Implementação de Ayres
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
		AplicarBonus(corrida)
		corrida.BonusAplicado = true
	} else if diferenca == 0 || diferenca > 0 && diferenca <= 15 {
		corrida.Status = models.StatusConcluidaNoTempo
		// Nenhuma penalização ou bônus
	}
}

func AvaliarCorrida(id int, nota int) error {
    for i := range corridas {
        if corridas[i].ID == id {
            corridas[i].Avaliacao = &nota
            return nil
        }
    }
    return errors.New("corrida não encontrada")
}

func (s *CorridaService) AdicionarCorrida(corrida models.Corrida) {
	corridas = append(corridas, corrida)
}

func CarregarCorridasDoArquivo() {
	file, err := os.Open("data/corridas.json")
	if err != nil {
		log.Println("Erro ao abrir arquivo JSON de corridas:", err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&corridas)
	if err != nil {
		log.Println("Erro ao fazer parse do JSON:", err)
	}
}

func GetCorridas() []models.Corrida {
	return corridas
}
