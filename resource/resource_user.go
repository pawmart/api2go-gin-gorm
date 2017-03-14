package resource

import (
	"errors"
	"net/http"

	"api2go-gin-gorm-simple/model"
	"api2go-gin-gorm-simple/storage"
	"github.com/manyminds/api2go"
)

// UserResource for api2go routes
type UserResource struct {
	UserStorage *storage.UserStorage
}

// FindAll to satisfy api2go data source interface
func (s UserResource) FindAll(r api2go.Request) (api2go.Responder, error) {
	users, err := s.UserStorage.GetAll()
	if err != nil {
		return &Response{}, err
	}

	return &Response{Res: users}, nil
}

// PaginatedFindAll can be used to load users in chunks
func (s UserResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	users, err := s.UserStorage.GetAll()
	if err != nil {
		return 0, &Response{}, err
	}

	// TODO: finish this off!
	return uint(len(users)), &Response{Res: users}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the user with the given ID, otherwise an error
func (s UserResource) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	user, err := s.UserStorage.GetOne(ID)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(err, err.Error(), http.StatusNotFound)
	}

	return &Response{Res: user}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (s UserResource) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	user, ok := obj.(model.User)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	id, err := s.UserStorage.Insert(user)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(errors.New("Faild to create a user"), "Faild to create a user", http.StatusInternalServerError)
	}
	err = user.SetID(id)
	if err != nil {
		return &Response{}, api2go.NewHTTPError(errors.New("Non-integer ID given"), "Non-integer ID given", http.StatusInternalServerError)
	}

	return &Response{Res: user, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (s UserResource) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	err := s.UserStorage.Delete(id)
	return &Response{Code: http.StatusNoContent}, err
}

//Update stores all changes on the user
func (s UserResource) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	user, ok := obj.(model.User)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New("Invalid instance given"), "Invalid instance given", http.StatusBadRequest)
	}

	err := s.UserStorage.Update(user)
	return &Response{Res: user, Code: http.StatusNoContent}, err
}
