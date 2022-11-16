package mysql_repository

import (
	"database/sql"

	e "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/entity"
)

type FornecedorRepository struct {
	DB *sql.DB
}

func NovoFornecedorRepository(db *sql.DB) *FornecedorRepository {
	return &FornecedorRepository{DB: db}
}

func (f *FornecedorRepository) Criar(fornecedor *e.Fornecedor) error {

	stmt, err := f.DB.Prepare("insert into fornecedores(id, razao_social, nome_fantasia, cnpj, ie, id_pecas, ativo, id_endereco, id_contato) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		fornecedor.GetID(),
		fornecedor.GetRazaoSocial(),
		fornecedor.GetNomeFantasia(),
		fornecedor.GetCNPJ(),
		fornecedor.GetCNPJ(),
		fornecedor.GetIdPecas(),
		fornecedor.IsAtivo(),
		1,
		1,
	)
	if err != nil {
		return err
	}

	return nil
}

// func (f *FornecedorRepository) BuscarPorID(id string) (*e.Fornecedor, error) {
// 	var fornecedor e.Fornecedor
// 	err := f.DB.First(&fornecedor, "id = ?", id).Error
// 	return &fornecedor, err
// }

// func (f *FornecedorRepository) BuscarTodos() ([]e.Fornecedor, error) {
// 	return nil, nil
// }

// func (f *FornecedorRepository) Atualizar(fornecedor *e.Fornecedor) error {
// 	return nil
// }

// func (f *FornecedorRepository) Excluir(id string) error {
// 	return nil
// }

// func (f *FornecedorRepository) BuscarPaginado(page, limit int, sort string) ([]e.Fornecedor, error) {
// 	var fornecedores []e.Fornecedor
// 	var err error
// 	if sort != "" && sort != "asc" && sort != "desc" {
// 		sort = "asc"
// 	}
// 	if page != 0 && limit != 0 {
// 		err = f.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&fornecedores).Error
// 	} else {
// 		err = f.DB.Order("created_at " + sort).Find(&fornecedores).Error
// 	}
// 	return fornecedores, err
// }
