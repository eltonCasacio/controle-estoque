package mysql_repository

import (
	"database/sql"

	"github.com/eltonCasacio/controle-estoque/internal/domain/peca/entity"
)

type PecaRepository struct {
	DB *sql.DB
}

func NovoPecaRepository(db *sql.DB) *PecaRepository {
	return &PecaRepository{DB: db}
}

func (f *PecaRepository) Criar(peca entity.Peca) error {
	f.DB.Prepare("insert into pecas(id, id_fornecedor, codigo, descricao, materiaprima, url_desenho_tecnico, url_foto, descricao_tecnica, massa, quantidade) values(?,?,?,?,?,?,?,?,?,?)")
	return nil
}
