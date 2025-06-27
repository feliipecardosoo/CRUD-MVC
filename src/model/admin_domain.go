package model

func NewAdminDomain(
	usuario string,
	senha string,
) AdminDomainInterface {
	return &adminDomain{
		usuario: usuario,
		senha:   senha,
	}
}

type AdminDomainInterface interface {
	GetUsuario() string
}

type adminDomain struct {
	usuario, senha string
}

func (d *adminDomain) GetUsuario() string {
	return d.usuario
}
