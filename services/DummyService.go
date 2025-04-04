package services

import (
	"your-app/database"
	"your-app/models"

	"gorm.io/gorm"
)

func ListDummyUser() ([]models.DummyUser, error) {
	var users []models.DummyUser
	err := database.GetDB().Find(&users).Error
	if err != nil {
		return []models.DummyUser{}, err
	}

	return users, nil
}

func GetDummyUser(id int) (models.DummyUser, error) {
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

func CreateDummyUser(user *models.DummyUser) error {
	err := database.GetDB().Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateDummyUser(id int, updateData *models.DummyUser) (models.DummyUser, error) {
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

func DeleteDummyUser(id int) error {
	err := database.GetDB().Delete(&models.DummyUser{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
