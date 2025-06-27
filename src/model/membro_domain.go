package model

import "crud/src/controller/model/request"

func NewMembroDomain(
	name string, dataNascimento string, anoBatismo int16,
	sexo string, estadoCivil string, dataCasamento string,
	nomeConjuge string, filho bool, email string,
	telefone string, status string, dataStatus string,
	endereco enderecoRequest,
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
		status:         status,
		dataStatus:     dataStatus,
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
	GetValidado() bool
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

func (d *membroDomain) GetValidado() bool {
	return d.validado
}
