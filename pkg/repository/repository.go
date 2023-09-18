package repository

import (
	"eff_mob_test/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type UserStorage interface {
	CreateUser(main_models.User) error
	CreateSingleUserRedis(user *main_models.User) error
	GetUser(int, int, map[string]interface{}) []main_models.User
	GetSingleUser(int) (*main_models.User, error)
	GetSingleUserRedis(int) (*JSONResponse, error)
	DeleteUser(int) error
	DeleteSingleUserRedis(int) error
	UpdateUser(main_models.User, []byte) (*main_models.User, error)
	UpdateSingleUserRedis(int, *main_models.User) error
	GetRedisConfig() *redis.Options
}

type Repository struct {
	UserStorage
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		NewUserStoragePostgres(db),
	}
}
