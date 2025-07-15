package services

import (
	"your-app/database"
	"your-app/models"
	"gorm.io/gorm"
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
