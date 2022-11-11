package mapper

import (
	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/webserver/handlers/usuario-handler/model"
)

func ConvertUsuarioDomainToModel(domain *entity.Usuario) *model.Usuario {
	return &model.Usuario{
		Id:    domain.GetID(),
		Nome:  domain.GetNome(),
		Senha: domain.GetSenha(),
		Ativo: domain.IsAtivo(),
	}
}
