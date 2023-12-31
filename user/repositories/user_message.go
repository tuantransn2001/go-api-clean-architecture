package repositories

import "food-comming-api/user/entities"

type UserMessaging interface {
	PushNotification(m *entities.UserPushNotificationDto) error
}
