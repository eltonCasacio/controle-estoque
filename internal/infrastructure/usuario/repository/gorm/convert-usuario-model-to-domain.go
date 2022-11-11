package gorm_repository

import "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"

func ConvertUsuarioModelToDomain(model *Usuario) *entity.Usuario {
	usuario := &entity.Usuario{}
	usuario.ChangeID(model.Id)
	usuario.ChangeNome(model.Nome)
	usuario.ChangeSenha(model.Senha)
	if model.Ativo {
		usuario.Ativar()
	} else {
		usuario.Desativar()
	}
	return usuario
}
