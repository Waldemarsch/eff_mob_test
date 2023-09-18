package repository

import (
	main_models "eff_mob_test/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserStoragePostgres struct {
	db *gorm.DB
}

func NewUserStoragePostgres(db *gorm.DB) *UserStoragePostgres {
	return &UserStoragePostgres{
		db: db,
	}
}

func (r *UserStoragePostgres) CreateUser(u main_models.User) error {
	result := r.db.Create(&u)

	err := result.Error

	if err != nil {
		return err
	}

	logrus.Infof("User %s %s with id %d is inserted into DB\n", u.Name, u.Surname, u.ID)

	return nil
}

func (r *UserStoragePostgres) GetUser(limit, page int, filters map[string]interface{}) []main_models.User {
	var users []main_models.User

	err := r.db.Scopes(NewPaginate(limit, page).PaginatedResult).Where(filters).Find(&users).Error

	if err != nil {
		logrus.Errorln(err)
		return nil
	}

	return users
}

func (r *UserStoragePostgres) GetSingleUser(id int) (*main_models.User, error) {
	var users *main_models.User

	err := r.db.Where(id).First(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserStoragePostgres) DeleteUser(id int) error {
	err := r.db.Delete(&main_models.User{ID: id}).Error

	if err != nil {
		return err
	}

	logrus.Infof("User %d is deleted\n", id)

	return nil
}

func (r *UserStoragePostgres) UpdateUser(u main_models.User, params []byte) (*main_models.User, error) {
	r.db.First(&u)

	err := json.Unmarshal(params, &u)

	if err != nil {
		return nil, err
	}

	r.db.Save(&u)

	logrus.Infof("User %d is updated\n", u.ID)

	return &u, nil
}
