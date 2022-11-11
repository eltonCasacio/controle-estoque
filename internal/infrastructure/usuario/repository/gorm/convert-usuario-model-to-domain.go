package gorm_repository

import "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"

func ConvertUsuarioModelToDomain(model *Usuario) *entity.Usuario {
	return &entity.Usuario{
		Id:    model.Id,
		Nome:  model.Nome,
		Senha: model.Senha,
		Ativo: model.Ativo,
	}
}
