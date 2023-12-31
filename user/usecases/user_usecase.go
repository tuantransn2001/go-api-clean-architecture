package usecases

import "food-comming-api/user/models"

type UserUseCase interface {
	UserDataProcessing(in *models.AddUserData) error
}
