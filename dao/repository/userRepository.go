package repository

import "github.com/claudioontheweb/go-dao-pattern/models"

type UserDao interface {
	Create(u *models.User) error
	GetById(i int) (models.User, error)
	GetAll() ([]models.User, error)
	Update(u *models.User) error
	Delete(i int) error
}