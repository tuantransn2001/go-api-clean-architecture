package repositories

import "food-comming-api/user/entities"

type UserRepository interface {
	InsertOne(in *entities.InsertUserDto) error
}
