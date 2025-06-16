package model

import (
	resterr "crud/src/configuration/rest-err"
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

type membroDomain struct {
	email    string
	password string
	name     string
	age      int8
}

type MembroDomainInterface interface {
	CreateMembroModel() *resterr.RestErr
	UpdateMembro(string) *resterr.RestErr
	FindMembro(string) *resterr.RestErr
	DeleteMembro(string) *resterr.RestErr
}

func (ud *membroDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
