package mapper

import (
	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/webserver/handlers/usuario-handler/model"
)

func ConvertUsuarioModelToDomain(model *model.Usuario) *entity.Usuario {
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
