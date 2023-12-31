package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	InsertUserDto struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}

	User struct {
		ID        *uuid.UUID `gorm:"type:char(36);primary_key"`
		Name      string     `gorm:"type:varchar(100);not null"`
		Email     string     `gorm:"type:varchar(100);uniqueIndex;not null"`
		Password  string     `gorm:"type:varchar(100);not null"`
		Role      *string    `gorm:"type:varchar(50);default:'user';not null"`
		Photo     *string    `gorm:"not null;default:'default.png'"`
		Verified  *bool      `gorm:"not null;default:false"`
		CreatedAt *time.Time `gorm:"not null;default:now()"`
		UpdatedAt *time.Time `gorm:"not null;default:now()"`
	}

	UserPushNotificationDto struct {
		Title        string `json:"title"`
		Body         string `json:"body"`
		ReportedTime string `json:"createdAt"`
	}
)

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	user.ID = &uuid
	return
}
