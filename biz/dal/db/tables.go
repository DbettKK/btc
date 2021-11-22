package db

import "gorm.io/gorm"

type TokenTable struct {
	gorm.Model
	Id    int64  `json:"id"`
	Token string `json:"token"`
}
