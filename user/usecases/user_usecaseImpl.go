package usecases

import (
	"food-comming-api/user/entities"
	"food-comming-api/user/models"
	"food-comming-api/user/repositories"
	"time"
)

type userUserCaseImpl struct {
	userRepository repositories.UserRepository
	userMessaging  repositories.UserMessaging
}

func NewUserUseCaseImpl(userRepository repositories.UserRepository, userMessaging repositories.UserMessaging) UserUseCase {
	return &userUserCaseImpl{
		userRepository: userRepository,
		userMessaging:  userMessaging,
	}
}

func (u *userUserCaseImpl) UserDataProcessing(in *models.AddUserData) error {
	insertUserData := &entities.InsertUserDto{
		Name:     in.Name,
		Email:    in.Email,
		Role:     in.Role,
		Password: in.Password,
	}

	if err := u.userRepository.InsertOne(insertUserData); err != nil {
		return err
	}

	pushUserData := &entities.UserPushNotificationDto{
		Title:        "New User",
		Body:         insertUserData.Name + " has been added",
		ReportedTime: time.Now().Local().Format("20023-12-31 20:20:20"),
	}

	if err := u.userMessaging.PushNotification(pushUserData); err != nil {
		return err
	}

	return nil
}
