package model

import (
	resterr "crud/src/configuration/rest-err"
	"crypto/md5"
	"encoding/hex"
)

func NewMembroDomain(email, password, name string, age int8) MembroDomainInterface {
	return &MembroDomain{
		email,
		password,
		name,
		age,
	}
}

type MembroDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *MembroDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type MembroDomainInterface interface {
	CreateMembro() *resterr.RestErr
	UpdateMembro(string) *resterr.RestErr
	FindMembro(string) *resterr.RestErr
	DeleteMembro(string) *resterr.RestErr
}
