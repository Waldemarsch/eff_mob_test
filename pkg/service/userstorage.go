package service

import (
	main_models "eff_mob_test/models"
	"eff_mob_test/pkg/repository"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

type UserStorageService struct {
	repo repository.UserStorage
}

func NewUserStorageService(repo repository.UserStorage) *UserStorageService {
	return &UserStorageService{
		repo: repo,
	}
}

func (s *UserStorageService) CreateUser(userJSON []byte) error {
	var msgJSON main_models.User
	err := json.Unmarshal(userJSON, &msgJSON)
	if err != nil {
		return err
	}
	if msgJSON.Name == "" {
		err = errors.New("error: Name is missing")
		return err
	}

	if msgJSON.Surname == "" {
		err = errors.New("error: surname is missing")
		return err
	}

	if msgJSON.Age == 0 {
		msgJSON.Age = GetAgeAPI("https://api.agify.io/", msgJSON.Name)
	}
	if msgJSON.Gender == "" {
		msgJSON.Gender = GetGenderAPI("https://api.genderize.io/", msgJSON.Name)
	}
	if msgJSON.Nationality == "" {
		msgJSON.Nationality = GetNationalityAPI("https://api.nationalize.io/", msgJSON.Name)
	}
	logrus.Infof("%s %s, Age: %d, Gender: %s, Nationality: %s\n",
		msgJSON.Name,
		msgJSON.Surname,
		msgJSON.Age,
		msgJSON.Gender,
		msgJSON.Nationality)

	err = s.repo.CreateUser(msgJSON)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStorageService) GetUser(params map[string][]string) ([]main_models.User, error) {
	limit, err := strconv.Atoi(params["limit"][0])

	delete(params, "limit")

	if err != nil {
		return nil, err
	}

	page, err := strconv.Atoi(params["page"][0])

	delete(params, "page")

	if err != nil {
		return nil, err
	}

	filters := make(map[string]interface{})

	for k, v := range params {
		filters[k] = v[0]
	}

	users := s.repo.GetUser(limit, page, filters)

	return users, nil
}

func (s *UserStorageService) GetSingleUser(id int) (*main_models.User, error) {
	redisResponse, err := s.repo.GetSingleUserRedis(id)

	if err != nil {
		logrus.Errorln("error while getting cache:", err)
		user, err := s.repo.GetSingleUser(id)

		if err != nil {
			return nil, err
		}

		err = s.repo.CreateSingleUserRedis(user)

		if err != nil {
			return nil, err
		}

		return user, err
	}

	return redisResponse.Data, nil
}

func (s *UserStorageService) DeleteUser(id int) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	err = s.repo.DeleteSingleUserRedis(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStorageService) UpdateUser(id int, params []byte) error {
	u := main_models.User{
		ID: id,
	}

	updUser, err := s.repo.UpdateUser(u, params)

	if err != nil {
		return err
	}

	err = s.repo.UpdateSingleUserRedis(id, updUser)

	if err != nil {
		return err
	}

	return nil
}
