package resource

import (
	"errors"
	"net/http"

	"api2go-gin-gorm-simple/model"
	"api2go-gin-gorm-simple/storage"
	"github.com/manyminds/api2go"
	"strconv"
)

// UserResource for api2go routes
type UserResource struct {
	UserStorage *storage.UserStorage
}

// FindAll to satisfy api2go data source interface
func (s UserResource) FindAll(r api2go.Request) (api2go.Responder, error) {

	users, _ := s.UserStorage.GetAll()

	var result []model.User

	for _, user := range users {
		result = append(result, *user)
	}

	return &Response{Res: result}, nil

}

// PaginatedFindAll can be used to load users in chunks
func (s UserResource) PaginatedFindAll(r api2go.Request) (uint, api2go.Responder, error) {
	users, err := s.UserStorage.GetAll()
	if err != nil {
		return 0, &Response{}, err
	}

	var (
		result       []model.User
		keys         []string
		number, size string
	)

	for k := range users {
		i := k
		if err != nil {
			return 0, &Response{}, err
		}

		keys = append(keys, i)
	}
	//sort.Ints(keys)

	numberQuery, ok := r.QueryParams["page[number]"]
	if ok {
		number = numberQuery[0]
	}
	sizeQuery, ok := r.QueryParams["page[size]"]
	if ok {
		size = sizeQuery[0]
	}

	if size != "" {
		sizeI, err := strconv.ParseUint(size, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		numberI, err := strconv.ParseUint(number, 10, 64)
		if err != nil {
			return 0, &Response{}, err
		}

		start := sizeI * (numberI - 1)
		for i := start; i < start + sizeI; i++ {
			if i >= uint64(len(users)) {
				break
			}
			result = append(result, *users[keys[i]])
		}
	}

	return uint(len(users)), &Response{Res: result}, nil
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
