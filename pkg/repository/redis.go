package repository

import (
	"context"
	main_models "eff_mob_test/models"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type JSONResponse struct {
	Data   *main_models.User `json:"data"`
	Source string            `json:"source"`
}

func (r *UserStoragePostgres) GetRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
}

func (r *UserStoragePostgres) CreateSingleUserRedis(user *main_models.User) error {

	redisClient := redis.NewClient(r.GetRedisConfig())

	ctx := context.Background()

	userJSON, err := json.Marshal(user)

	if err != nil {
		return err
	}

	status := redisClient.Set(ctx, fmt.Sprintf("%d", user.ID), userJSON, time.Minute*10)

	_, err = status.Result()

	if err != nil {
		return err
	}

	logrus.Infof("User with id %d is added to redis\n", user.ID)

	return nil

}

func (r *UserStoragePostgres) GetSingleUserRedis(id int) (*JSONResponse, error) {

	redisClient := redis.NewClient(r.GetRedisConfig())

	ctx := context.Background()

	cachedUser, err := redisClient.Get(ctx, fmt.Sprintf("%d", id)).Bytes()

	if err != nil {
		return nil, err
	}

	logrus.Infof("found user %d in redis\n", id)

	var user main_models.User

	err = json.Unmarshal(cachedUser, &user)

	if err != nil {
		return nil, err
	}

	response := JSONResponse{Data: &user, Source: "Redis Cache"}

	return &response, nil

}

func (r *UserStoragePostgres) DeleteSingleUserRedis(id int) error {
	redisClient := redis.NewClient(r.GetRedisConfig())

	ctx := context.Background()

	status := redisClient.Del(ctx, fmt.Sprintf("%d", id))

	_, err := status.Result()

	if err != nil {
		return err
	}

	logrus.Infof("User %d is deleted\n", id)

	return nil
}

func (r *UserStoragePostgres) UpdateSingleUserRedis(id int, user *main_models.User) error {
	redisClient := redis.NewClient(r.GetRedisConfig())

	ctx := context.Background()

	statusDelete := redisClient.Del(ctx, fmt.Sprintf("%d", id))

	_, err := statusDelete.Result()

	if err != nil {
		return err
	}

	userJSON, err := json.Marshal(user)

	if err != nil {
		return err
	}

	statusCreate := redisClient.Set(ctx, fmt.Sprintf("%d", id), userJSON, time.Minute*10)

	_, err = statusCreate.Result()

	if err != nil {
		return err
	}

	logrus.Infof("User %d is updated\n", id)

	return nil
}
