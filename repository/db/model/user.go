package model

import (
	"github.com/CocaineCong/secret"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	conf "test-gin-mall/config"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          string
	Relation       []User `gorm:"many2many:relation;"`
}

const (
	Active       string = "active"
	PasswordCost        = 12
)

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

func (u *User) EncryptMoney(key string) (money string, err error) {
	aseEncrypt, err := secret.NewAesEncrypt(conf.Config.Encrypt.MoneyEncrypt, key, "", secret.AesEncrypt128, secret.AesModeTypeCBC)

	if err != nil {
		return
	}
	money = aseEncrypt.SecretEncrypt(u.Money)

	return money, nil
}

func (u *User) DecryptMoney(key string) (money float64, err error) {
	aseDecrypt, err := secret.NewAesEncrypt(conf.Config.Encrypt.MoneyEncrypt, key, "", secret.AesEncrypt128, secret.AesModeTypeCBC)
	if err != nil {
		return
	}
	money = cast.ToFloat64(aseDecrypt.SecretDecrypt(u.Money))
	return
}
