package gorm_repository

import "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"

func ConvertUsuarioDomainToModel(domain *entity.Usuario) *Usuario {
	return &Usuario{
		Id:    domain.GetID(),
		Nome:  domain.GetNome(),
		Senha: domain.GetSenha(),
		Ativo: domain.IsAtivo(),
	}
}
