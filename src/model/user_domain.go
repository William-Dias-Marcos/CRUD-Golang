package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface{
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	
	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string 
	age      int8
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

