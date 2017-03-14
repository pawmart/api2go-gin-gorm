package storage

import (
	"errors"
	"fmt"
	"net/http"

	"api2go-gin-gorm-simple/model"
	"github.com/jinzhu/gorm"
	"github.com/manyminds/api2go"
)

// NewUserStorage initializes the storage
func NewUserStorage(db *gorm.DB) *UserStorage {
	return &UserStorage{db}
}

// UserStorage stores all users
type UserStorage struct {
	db *gorm.DB
}

// GetAll returns the user map (because we need the ID as key too)
func (s UserStorage) GetAll() (map[string]*model.User, error) {

	var users []model.User
	s.db.Find(&users)
	if s.db.Error != nil {
		return nil, s.db.Error
	}

	userMap := make(map[string]*model.User)
	for i := range users {
		u := &users[i]
		userMap[u.ID] = u
	}
	return userMap, nil
}

// GetOne user
func (s UserStorage) GetOne(id string) (model.User, error) {

	var user model.User
	s.db.First(&user, "id = ?", id)

	if err := s.db.Error; err == gorm.ErrRecordNotFound {
		errMessage := fmt.Sprintf("User for id %s not found", id)
		return model.User{}, api2go.NewHTTPError(errors.New(errMessage), errMessage, http.StatusNotFound)
	} else if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// Insert a user
func (s *UserStorage) Insert(c model.User) (string, error) {

	s.db.Create(&c)
	if s.db.Error != nil {
		return "", s.db.Error
	}
	return c.GetID(), nil
}

// Delete one :(
func (s *UserStorage) Delete(id string) error {

	var user model.User
	s.db.First(&user, "id = ?", id)
	if err := s.db.Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("User with id %s does not exist", id)
	}
	s.db.Delete(&user)

	return s.db.Error
}

// Update a user
func (s *UserStorage) Update(c model.User) error {

	s.db.Save(&c)
	return s.db.Error
}