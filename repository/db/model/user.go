package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Avatar         string `gorm:"size:1000"`
	Money          string
	Relation       []User `gorm:"many2many:relation;"`
}
