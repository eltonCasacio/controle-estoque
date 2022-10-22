package database

import (
	e "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/entity"
	"gorm.io/gorm"
)

type FornecedorRepository struct {
	DB *gorm.DB
}

func NovoFornecedorRepository(db *gorm.DB) *FornecedorRepository {
	return &FornecedorRepository{DB: db}
}

func (f *FornecedorRepository) Criar(fornecedor *e.Fornecedor) error {
	return f.DB.Create(fornecedor).Error
}

func (f *FornecedorRepository) BuscarPorID(id int64) (*e.Fornecedor, error) {
	var fornecedor e.Fornecedor
	err := f.DB.First(&fornecedor, "id = ?", id).Error
	return &fornecedor, err
}

func (f *FornecedorRepository) BuscarTodos(fornecedor e.Fornecedor) ([]e.Fornecedor, error) {
	return nil, nil
}

func (f *FornecedorRepository) Atualizar(fornecedor e.Fornecedor) error {
	return nil
}

func (f *FornecedorRepository) Excluir(id int64) error {
	return nil
}

func (f *FornecedorRepository) BuscarPaginado(page, limit int, sort string) ([]e.Fornecedor, error) {
	var fornecedores []e.Fornecedor
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = f.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&fornecedores).Error
	} else {
		err = f.DB.Order("created_at " + sort).Find(&fornecedores).Error
	}
	return fornecedores, err
}
