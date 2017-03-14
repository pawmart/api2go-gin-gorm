package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lonelycode/go-uuid/uuid"
)

// User is a generic database user
type User struct {
	ID string `json:"-"`
	//rename the username field to user-name.
	Username      string      `json:"user-name"`
	PasswordHash  string      `json:"-"`
	exists        bool        `sql:"-"`
}

// Generate uuid not int id.
func (user *User) BeforeCreate(scope *gorm.Scope) error {
  scope.SetColumn("ID", uuid.New())
  return nil
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
