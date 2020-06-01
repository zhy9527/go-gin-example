package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Auth struct {
	Model

	Username   string `json:"username"`
	Password   string `json:"password"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

func ExistAuthByUsername(username string) bool {
	var auth Auth
	db.Where("username = ?", username).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

func AddAuth(username string, password string, created_by string) bool {
	db.Create(&Auth{
		Username:  username,
		Password:  password,
		CreatedBy: created_by,
	})

	return true
}

func (auth *Auth) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (auth *Auth) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
