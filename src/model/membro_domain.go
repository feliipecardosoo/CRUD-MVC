package model

import (
	"crypto/md5"
	"encoding/hex"
)

func NewMembroDomain(email, password, name string, age int8) MembroDomainInterface {
	return &membroDomain{
		email,
		password,
		name,
		age,
	}
}

type MembroDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() string
	GetName() string

	EncryptPassword()
}

type membroDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (d *membroDomain) GetEmail() string {
	return d.email
}

func (d *membroDomain) GetPassword() string {
	return d.password
}

func (d *membroDomain) GetName() string {
	return d.name
}

func (d *membroDomain) GetAge() int8 {
	return d.age
}

func (ud *membroDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
