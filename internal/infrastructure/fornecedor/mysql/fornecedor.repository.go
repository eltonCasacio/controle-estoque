package mysql_repository

import (
	"database/sql"

	e "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/entity"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/fornecedor/model"
)

type FornecedorRepository struct {
	DB *sql.DB
}

func NovoFornecedorRepository(db *sql.DB) *FornecedorRepository {
	return &FornecedorRepository{DB: db}
}

func (f *FornecedorRepository) Criar(fornecedor *e.Fornecedor) error {
	tx, err := f.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("insert into fornecedores(id, razao_social, nome_fantasia, cnpj, ie, ativo) values(?,?,?,?,?,?)")
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

	stmt, err = tx.Prepare("insert into fornecedores_pecas(id_fornecedor, id_peca) values(?,?)")
	if err != nil {
		return err
	}

	for _, v := range fornecedor.GetIdPecas() {
		_, err = stmt.Exec(fornecedor.GetID(), v)
		if err != nil {
			return err
		}
	}

	stmt, err = tx.Prepare("insert into enderecos(uf, rua, complemento, bairro, cep, numero, fornecedor_id) values(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		fornecedor.GetEndereco().UF,
		fornecedor.GetEndereco().Rua,
		fornecedor.GetEndereco().Complemento,
		fornecedor.GetEndereco().Bairro,
		fornecedor.GetEndereco().CEP,
		fornecedor.GetEndereco().Numero,
		fornecedor.GetID(),
	)
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("insert into contatos(email, celular, nome, fornecedor_id) values(?,?,?,?)")
	if err != nil {
		return err
	}

	for _, contato := range fornecedor.GetContatos() {
		_, err = stmt.Exec(
			contato.Email,
			contato.Celular,
			contato.Nome,
			fornecedor.GetID(),
		)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (f *FornecedorRepository) BuscarPorID(id string) (*e.Fornecedor, error) {
	stmt, err := f.DB.Prepare("select * from fornecedores where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var fornecedorModel model.FornecedorModel
	err = stmt.QueryRow(id).Scan(&fornecedorModel.Id, &fornecedorModel.RazaoSocial, &fornecedorModel.NomeFantasia, &fornecedorModel.CNPJ, &fornecedorModel.Ie, &fornecedorModel.Ativo)
	if err != nil {
		return nil, err
	}

	var fornecedor e.Fornecedor
	fornecedor.ChangeID(fornecedorModel.Id)
	fornecedor.ChangeNomeFantasia(fornecedorModel.NomeFantasia)
	fornecedor.ChangeRazaoSocial(fornecedorModel.RazaoSocial)
	fornecedor.ChangeCNPJ(fornecedorModel.CNPJ)
	fornecedor.ChangeIe(fornecedorModel.Ie)
	fornecedor.Desativar()
	if fornecedorModel.Ativo {
		fornecedor.Ativar()
	}
	return &fornecedor, nil
}

func (f *FornecedorRepository) BuscarTodos() ([]e.Fornecedor, error) {
	rows, err := f.DB.Query("select * from fornecedores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fornecedores []e.Fornecedor

	for rows.Next() {
		var fornecedor model.FornecedorModel
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

func (f *FornecedorRepository) Atualizar(fornecedor *e.Fornecedor) error {
	stmt, err := f.DB.Prepare("update fornecedores set id = ?, razao_social = ?, nome_fantasia = ?, cnpj = ?, ie = ?, ativo = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(fornecedor.GetID(), fornecedor.GetRazaoSocial(), fornecedor.GetNomeFantasia(), fornecedor.GetCNPJ(), fornecedor.GetIe(), fornecedor.IsAtivo())
	if err != nil {
		return err
	}
	return nil
}

func (f *FornecedorRepository) Excluir(id string) error {
	stmt, err := f.DB.Prepare("delete from fornecedores where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (f *FornecedorRepository) BuscarPaginado(page, limit int, sort string) ([]e.Fornecedor, error) {
	rows, err := f.DB.Query("SELECT * from fornecedores LIMIT ?, ?;", page, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fornecedores []e.Fornecedor

	for rows.Next() {
		var fornecedor model.FornecedorModel
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
