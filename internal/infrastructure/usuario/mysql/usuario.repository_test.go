package gorm_repository

import (
	"database/sql"
	"testing"

	usuario_entity "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	model "github.com/eltonCasacio/controle-estoque/internal/infrastructure/usuario/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func Repository() (*sql.DB, *UsuarioRepository, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/safisa")
	repo := NovoUsuarioRpository(db)

	stmt, _ := db.Prepare("delete from usuarios")
	stmt.Exec()
	defer stmt.Close()
	return db, repo, err
}

func TestCriarUsuario(t *testing.T) {
	db, repository, err := Repository()
	assert.Nil(t, err)
	usuario, _ := usuario_entity.NovoUsuario("roberto", "123")
	err = repository.Criar(usuario)
	assert.Nil(t, err)

	stmt, err := db.Prepare("select id, nome from usuarios where id = ?")
	if err != nil {
		panic(err)
	}
	var usuarioModel model.UsuarioModel
	err = stmt.QueryRow(usuario.GetID()).Scan(&usuarioModel.Id, &usuarioModel.Nome)
	assert.Nil(t, err)
	assert.Equal(t, usuario.GetID(), usuarioModel.Id)
}

func TestBuscarUsuarioPorId(t *testing.T) {
	_, repository, _ := Repository()
	usuario, _ := usuario_entity.NovoUsuario("roberto", "123")
	err := repository.Criar(usuario)
	assert.Nil(t, err)

	usuarioModel, err := repository.BuscarPorID(usuario.GetID().String())
	assert.Nil(t, err)
	assert.NotNil(t, usuarioModel)
	assert.Equal(t, usuario.GetID(), usuarioModel.GetID())
	assert.Equal(t, usuario.GetNome(), usuarioModel.GetNome())
}

func TestBuscarUsuario_QueNaoExiste(t *testing.T) {
	db, _, err := Repository()
	assert.Nil(t, err)
	stmt, err := db.Prepare("select * from usuarios where id = ?")
	if err != nil {
		panic(err)
	}
	var usuarioModel model.UsuarioModel
	err = stmt.QueryRow("836gfe").Scan(&usuarioModel.Id, &usuarioModel.Nome, &usuarioModel.Senha, &usuarioModel.Ativo)
	assert.NotNil(t, err)
}

func TestBuscarTodos(t *testing.T) {
	_, repository, err := Repository()
	assert.Nil(t, err)

	usuario, _ := usuario_entity.NovoUsuario("Elton", "1223")
	_ = repository.Criar(usuario)
	usuario, _ = usuario_entity.NovoUsuario("Casacio", "12236")
	_ = repository.Criar(usuario)

	usuarios, err := repository.BuscarTodos()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(usuarios))
}

func TestAtualizarUsuario(t *testing.T) {
	_, repository, err := Repository()
	assert.Nil(t, err)
	usuario, _ := usuario_entity.NovoUsuario("Elton", "123")
	_ = repository.Criar(usuario)

	uEncontrado, _ := repository.BuscarPorID(usuario.GetID().String())
	uEncontrado.ChangeNome("Casacio")

	repository.Atualizar(uEncontrado)
	usuarioAtualizado, err := repository.BuscarPorID(usuario.GetID().String())
	assert.Nil(t, err)
	assert.NotNil(t, usuarioAtualizado)
	assert.Equal(t, "Casacio", usuarioAtualizado.GetNome())
}

func TestExcluirUsuario(t *testing.T) {
	_, repository, _ := Repository()
	usuario, _ := usuario_entity.NovoUsuario("delete_user", "123")
	_ = repository.Criar(usuario)

	err := repository.Excluir(usuario.GetID().String())
	assert.Nil(t, err)

	uEncontrado, err := repository.BuscarPorID(usuario.GetID().String())
	assert.NotNil(t, err)
	assert.Nil(t, uEncontrado)
}

func TestBuscarUsuarioPorNome(t *testing.T) {
	_, repository, _ := Repository()
	usuario1, _ := usuario_entity.NovoUsuario("usuario 1", "123")
	_ = repository.Criar(usuario1)

	usuario2, _ := usuario_entity.NovoUsuario("usuario 2", "123")
	_ = repository.Criar(usuario2)

	usuario3, _ := usuario_entity.NovoUsuario("usuario 3", "123")
	_ = repository.Criar(usuario3)

	usuario4, _ := usuario_entity.NovoUsuario("usuario 4", "123")
	_ = repository.Criar(usuario4)

	usuariosEncontrado, err := repository.BuscarPaginado(1, 3, "")
	assert.Nil(t, err)
	assert.NotNil(t, usuariosEncontrado)
	assert.Equal(t, 3, len(usuariosEncontrado))
}
