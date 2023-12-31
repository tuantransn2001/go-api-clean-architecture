package repositories

import (
	"food-comming-api/user/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) UserRepository {
	return &userPostgresRepository{
		db: db,
	}
}

func (u *userPostgresRepository) InsertOne(in *entities.InsertUserDto) error {
	result := u.db.Create(in)
	if result.Error != nil {
		log.Errorf("Insert one user: %s", result.Error)
		return result.Error
	}
	log.Debugf("Insert one user: %s", result.RowsAffected)
	return nil
}
