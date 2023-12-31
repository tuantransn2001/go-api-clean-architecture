package repositories

import (
	"food-comming-api/user/entities"

	"github.com/labstack/gommon/log"
)

type userFCMRepository struct{}

func NewUserFCMessaging() UserMessaging {
	return &userFCMRepository{}
}

func (u *userFCMRepository) PushNotification(m *entities.UserPushNotificationDto) error {
	// ? Push notification logic here...
	log.Debugf("Pushed FCM notification with data: %v", "m")
	return nil
}
