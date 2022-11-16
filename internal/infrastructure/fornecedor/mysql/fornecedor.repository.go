package mysql_repository

import (
	"database/sql"

	e "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/entity"
	fornecedor "github.com/eltonCasacio/controle-estoque/internal/infrastructure/fornecedor/model"
)

type FornecedorRepository struct {
	DB *sql.DB
}

func NovoFornecedorRepository(db *sql.DB) *FornecedorRepository {
	return &FornecedorRepository{DB: db}
}

func (f *FornecedorRepository) Criar(fornecedor *e.Fornecedor) error {

	stmt, err := f.DB.Prepare("insert into fornecedores(id, razao_social, nome_fantasia, cnpj, ie, ativo) values(?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		fornecedor.GetID(),
		fornecedor.GetRazaoSocial(),
		fornecedor.GetNomeFantasia(),
		fornecedor.GetCNPJ(),
		fornecedor.GetIe(),
		fornecedor.IsAtivo(),
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

func (f *FornecedorRepository) BuscarTodos() ([]e.Fornecedor, error) {
	rows, err := f.DB.Query("select * from fornecedores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fornecedores []e.Fornecedor

	for rows.Next() {
		var fornecedor fornecedor.FornecedorModel
		err := rows.Scan(&fornecedor.Id, &fornecedor.RazaoSocial, &fornecedor.NomeFantasia, &fornecedor.CNPJ, &fornecedor.Ie, &fornecedor.Ativo)
		if err != nil {
			return nil, err
		}
		var f e.Fornecedor
		f.ChangeID(fornecedor.Id)
		f.ChangeRazaoSocial(fornecedor.RazaoSocial)
		f.ChangeNomeFantasia(fornecedor.NomeFantasia)
		f.ChangeCNPJ(fornecedor.CNPJ)
		f.ChangeIe(fornecedor.Ie)
		f.Desativar()
		if fornecedor.Ativo {
			f.Ativar()
		}

		fornecedores = append(fornecedores, f)
	}
	return fornecedores, nil
}

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
