package model

import (
	"github.com/jinzhu/gorm"
	"api2go-gin-gorm-simple/utils"
)

// User is a generic database user
type User struct {
	ID           string `json:"-"`
	//rename the username field to user-name.
	Username     string      `json:"user-name"`
	PasswordHash string      `json:"-"`
	exists       bool        `sql:"-"`
}

// Generate human id
func (user *User) BeforeCreate(scope *gorm.Scope) error {

	identifier := utils.GenerateHumanId("USR")

	if scope.DB().Select("1").Where("id = ?", identifier).First(&user).RecordNotFound() == false {
		return user.BeforeCreate(scope)
	}

	return scope.SetColumn("ID", identifier)
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u User) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *User) SetID(id string) error {
	u.ID = id
	return nil
}
