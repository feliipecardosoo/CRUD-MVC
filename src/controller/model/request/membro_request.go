package request

type MembroRequest struct {
	Name           string          `json:"name" binding:"required,min=3,max=100"`
	DataNascimento string          `json:"data_nascimento" binding:"required,datetime=2006-01-02"`
	AnoBatismo     int16           `json:"ano_batismo" binding:"omitempty,min=1900,max=2100"`
	Sexo           string          `json:"sexo" binding:"required,oneof=masculino feminino"`
	EstadoCivil    string          `json:"estado_civil" binding:"required,oneof=solteiro casado divorciado viuvo"`
	DataCasamento  string          `json:"data_casamento" binding:"omitempty,datetime=2006-01-02"`
	NomeConjuge    string          `json:"nome_conjuge" binding:"omitempty,min=3,max=100"`
	Filho          bool            `json:"filho" binding:"required"`
	Email          string          `json:"email" binding:"required,email"`
	Telefone       string          `json:"telefone" binding:"omitempty,min=10,max=15"`
	Status         string          `json:"status" binding:"required,oneof=ativo inativo"`
	DataStatus     string          `json:"data_status" binding:"required,datetime=2006-01-02"`
	Endereco       EnderecoRequest `json:"endereco" binding:"required"`
}
