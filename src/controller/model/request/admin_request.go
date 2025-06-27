package request

type AdminRequest struct {
	Id      int64  `json:"id" binding:"omitempty"`
	Usuario string `json:"usuario" binding:"required"`
	Senha   string `json:"senha" binding:"required,min=6,max=100"`
}
