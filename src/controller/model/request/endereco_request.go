package request

type EnderecoRequest struct {
	Cep         string `json:"cep" binding:"required,min=8,max=9"`
	Rua         string `json:"rua" binding:"required,min=3,max=100"`
	Numero      string `json:"numero" binding:"required,min=1,max=10"`
	Bairro      string `json:"bairro" binding:"required,min=3,max=100"`
	Complemento string `json:"complemento" binding:"omitempty,min=0,max=100"`
}
