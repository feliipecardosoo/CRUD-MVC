package model

import (
	"crud/src/controller/model/request"
	"time"
)

func NewMembroDomain(
	name string, dataNascimento string, anoBatismo int16,
	sexo string, estadoCivil string, dataCasamento string,
	nomeConjuge string, filho bool, email string,
	telefone string, endereco enderecoRequest,
) MembroDomainInterface {
	return &membroDomain{
		name:           name,
		dataNascimento: dataNascimento,
		anoBatismo:     anoBatismo,
		sexo:           sexo,
		estadoCivil:    estadoCivil,
		dataCasamento:  dataCasamento,
		nomeConjuge:    nomeConjuge,
		filho:          filho,
		email:          email,
		telefone:       telefone,
		status:         "ativo",
		dataStatus:     time.Now().Format("2006-01-02"),
		endereco:       endereco,
		validado:       false,
	}
}

func ConvertEnderecoRequest(req request.EnderecoRequest) enderecoRequest {
	return enderecoRequest{
		cep:         req.Cep,
		rua:         req.Rua,
		numero:      req.Numero,
		bairro:      req.Bairro,
		complemento: req.Complemento,
	}
}

type MembroDomainInterface interface {
	GetName() string
	GetStatus() string
	GetDataStatus() string
}

type membroDomain struct {
	name           string
	dataNascimento string
	anoBatismo     int16
	sexo           string
	estadoCivil    string
	dataCasamento  string
	nomeConjuge    string
	filho          bool
	email          string
	telefone       string
	status         string
	dataStatus     string
	endereco       enderecoRequest
	validado       bool
}

type enderecoRequest struct {
	cep         string
	rua         string
	numero      string
	bairro      string
	complemento string
}

func (d *membroDomain) GetName() string {
	return d.name
}

func (d *membroDomain) GetStatus() string {
	return d.status
}

func (d *membroDomain) GetDataStatus() string {
	return d.dataStatus
}
