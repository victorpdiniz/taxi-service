package services

import (
	"time"
	"your-app/database"
	"your-app/models"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

func ListCorridas() ([]models.Corrida, error) {
	var corridas []models.Corrida
	err := database.GetDB().Find(&corridas).Error
	if err != nil {
		return []models.Corrida{}, err
	}

	return corridas, nil
}

func GetCorrida(id int) (models.Corrida, error) {
	var user models.DummyUser
	err := database.GetDB().First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.DummyUser{}, err
		}
		return models.DummyUser{}, err
	}
	return user, nil
}

func CreateCorrida(user *models.DummyUser) error {
	err := database.GetDB().Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateCorrida(id int, updateData *models.DummyUser) (models.DummyUser, error) {
	user, err := GetDummyUser(id)
	if err != nil {
		return models.DummyUser{}, err
	}

	// Update the user fields with the new data
	if updateData.Name != "" {
		user.Name = updateData.Name
	}
	if updateData.Email != "" {
		user.Email = updateData.Email
	}
	// Add other fields as needed

	err = database.GetDB().Model(&models.DummyUser{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return models.DummyUser{}, err
	}

	// Fetch the updated user to return
	updatedUser, err := GetDummyUser(id)
	if err != nil {
		return models.DummyUser{}, err
	}

	return updatedUser, nil
}

func DeleteCorrida(id int) error {
	err := database.GetDB().Delete(&models.DummyUser{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

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
