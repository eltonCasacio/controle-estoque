package gorm_repository

import "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"

func ConvertUsuarioDomainToModel(domain *entity.Usuario) *Usuario {
	return &Usuario{
		Id:    domain.Id,
		Nome:  domain.Nome,
		Senha: domain.Senha,
		Ativo: domain.Ativo || true,
	}
}
