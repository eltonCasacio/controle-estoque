package mysql_repository

import (
	"database/sql"
	"testing"

	usuario_entity "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"github.com/eltonCasacio/controle-estoque/internal/infrastructure/usuario/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UsuarioTestSuite struct {
	suite.Suite
	DB         *sql.DB
	Usuario    usuario_entity.Usuario
	Repository *UsuarioRepository
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UsuarioTestSuite))
}

func (suite *UsuarioTestSuite) SetupTest() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/safisa")
	if err != nil {
		panic(err)
	}
	suite.DB = db
	suite.Repository = NovoUsuarioRpository(db)

	usuario, _ := usuario_entity.NovoUsuario("roberto", "123")
	suite.Usuario = *usuario
	stmt, _ := suite.DB.Prepare("delete from usuarios")
	stmt.Exec()
}

func (suite *UsuarioTestSuite) TestCriarUsuario() {
	defer suite.DB.Close()
	err := suite.Repository.Criar(&suite.Usuario)
	assert.Nil(suite.T(), err)

	stmt, err := suite.DB.Prepare("select id, nome from usuarios where id = ?")
	if err != nil {
		panic(err)
	}

	var usuarioModel model.UsuarioModel
	err = stmt.QueryRow(suite.Usuario.GetID()).Scan(&usuarioModel.Id, &usuarioModel.Nome)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.Usuario.GetID(), usuarioModel.Id)
}

func (suite *UsuarioTestSuite) TestBuscarUsuarioPorId() {
	defer suite.DB.Close()
	err := suite.Repository.Criar(&suite.Usuario)
	assert.Nil(suite.T(), err)

	usuarioModel, err := suite.Repository.BuscarPorID(suite.Usuario.GetID().String())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), usuarioModel)
	assert.Equal(suite.T(), suite.Usuario.GetID(), usuarioModel.GetID())
	assert.Equal(suite.T(), suite.Usuario.GetNome(), usuarioModel.GetNome())
}

func (suite *UsuarioTestSuite) TestBuscarUsuario_QueNaoExiste() {
	defer suite.DB.Close()
	stmt, err := suite.DB.Prepare("select * from usuarios where id = ?")
	if err != nil {
		panic(err)
	}
	var usuarioModel model.UsuarioModel
	err = stmt.QueryRow("836gfe").Scan(&usuarioModel.Id, &usuarioModel.Nome, &usuarioModel.Senha, &usuarioModel.Ativo)
	assert.NotNil(suite.T(), err)
}

func (suite *UsuarioTestSuite) TestBuscarTodos() {
	defer suite.DB.Close()
	suite.Repository.Criar(&suite.Usuario)
	suite.Repository.Criar(&suite.Usuario)

	usuarios, err := suite.Repository.BuscarTodos()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(usuarios))
}

func (suite *UsuarioTestSuite) TestAtualizarUsuario() {
	defer suite.DB.Close()
	suite.Repository.Criar(&suite.Usuario)

	uEncontrado, _ := suite.Repository.BuscarPorID(suite.Usuario.GetID().String())
	uEncontrado.ChangeNome("Casacio")

	suite.Repository.Atualizar(uEncontrado)
	usuarioAtualizado, err := suite.Repository.BuscarPorID(suite.Usuario.GetID().String())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), usuarioAtualizado)
	assert.Equal(suite.T(), "Casacio", usuarioAtualizado.GetNome())
}

func (suite *UsuarioTestSuite) TestExcluirUsuario() {
	defer suite.DB.Close()
	suite.Repository.Criar(&suite.Usuario)

	err := suite.Repository.Excluir(suite.Usuario.GetID().String())
	assert.Nil(suite.T(), err)

	uEncontrado, err := suite.Repository.BuscarPorID(suite.Usuario.GetID().String())
	assert.NotNil(suite.T(), err)
	assert.Nil(suite.T(), uEncontrado)
}

func (suite *UsuarioTestSuite) TestBuscarPaginado() {
	defer suite.DB.Close()
	suite.Repository.Criar(&suite.Usuario)
	suite.Repository.Criar(&suite.Usuario)
	suite.Repository.Criar(&suite.Usuario)
	suite.Repository.Criar(&suite.Usuario)

	usuariosEncontrado, err := suite.Repository.BuscarPaginado("1", "3", "")
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), usuariosEncontrado)
	assert.Equal(suite.T(), 3, len(usuariosEncontrado))
}
