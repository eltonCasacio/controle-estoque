package mysql_repository

import (
	"database/sql"
	"testing"

	f_entity "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/entity"
	value_object "github.com/eltonCasacio/controle-estoque/internal/domain/fornecedor/value-object"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FornecedorTestSuite struct {
	suite.Suite
	Endereco   value_object.Endereco
	Contatos   []value_object.Contato
	Fornecedor f_entity.Fornecedor
	DB         *sql.DB
	Repository *FornecedorRepository
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(FornecedorTestSuite))
}

func (suite *FornecedorTestSuite) SetupTest() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/safisa")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	suite.DB = db
	suite.Repository = NovoFornecedorRepository(suite.DB)

	var contatos []value_object.Contato
	contatos = append(contatos, value_object.Contato{Nome: "elton", Telefone: "123456789", Email: "elton@mail.com", Celular: "1232423432"})
	contatos = append(contatos, value_object.Contato{Nome: "Roberto", Telefone: "123456789", Email: "roberto@mail.com", Celular: "1232423432"})
	contatos = append(contatos, value_object.Contato{Nome: "Daniel", Telefone: "123456789", Email: "daniel@mail.com", Celular: "1232423432"})
	suite.Contatos = contatos

	suite.Endereco = value_object.Endereco{Cidade: "valinhos", UF: "sp", Rua: "any_rua", Complemento: "any_complemento", Bairro: "any_bairro", CEP: 23423, Numero: "2345345"}
	fornecedor, _ := f_entity.NovoFornecedor("nome fantasia", suite.Endereco, suite.Contatos, []string{"1", "2"})
	suite.Fornecedor = *fornecedor
}

func (suite *FornecedorTestSuite) TestFornecedor_Criar() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/safisa")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repo := NovoFornecedorRepository(db)

	err = repo.Criar(&suite.Fornecedor)
	assert.Nil(suite.T(), err)

	// var fornecedor f_entity.Fornecedor
	// err = suite.DB.First(&fornecedor, "id = ?", suite.Fornecedor.GetID()).Error
	// assert.Nil(suite.T(), err)
}
