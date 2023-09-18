package service

import (
	"eff_mob_test/models"
	"eff_mob_test/pkg/repository"
)

type UserStorage interface {
	CreateUser([]byte) error
	GetUser(map[string][]string) ([]main_models.User, error)
	GetSingleUser(int) (*main_models.User, error)
	DeleteUser(int) error
	UpdateUser(int, []byte) error
}

type Service struct {
	UserStorage
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserStorage: NewUserStorageService(repo),
	}
}
