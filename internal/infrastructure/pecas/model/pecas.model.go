package model

import "github.com/eltonCasacio/controle-estoque/pkg/entity"

type PecaModel struct {
	Id                entity.ID
	IdFornecedor      string
	Codigo            string
	Descricao         string
	MateriaPrima      string
	UrlDesenhoTecnico string
	UrlFoto           string
	DescricaoTecnica  string
	Massa             float64
	Quantidade        int
}
