package fornecedor_repository

import (
	e "github.com/eltonCasacio/controle-estoque/src/domain/fornecedor/entity"
)

type FornecedorRepository struct {
}

func (f *FornecedorRepository) Create(fornecedor e.FornecedorInterface) error {
	return nil
}

func (f *FornecedorRepository) Find(id int64) (e.FornecedorInterface, error) {
	return nil, nil
}

func (f *FornecedorRepository) FindAll(fornecedor e.FornecedorInterface) ([]e.FornecedorInterface, error) {
	return nil, nil
}

func (f *FornecedorRepository) Update(fornecedor e.FornecedorInterface) error {
	return nil
}

func (f *FornecedorRepository) Delete(id int64) error {
	return nil
}
