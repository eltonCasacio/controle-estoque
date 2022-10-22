package database

import (
	"testing"

	usuario_entity "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Repository() (*gorm.DB, *UsuarioRepository, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&usuario_entity.Usuario{})
	repo := NovoUsuarioRpository(db)
	return db, repo, err
}

func TestCrriarUsuario(t *testing.T) {
	db, repository, err := Repository()
	assert.Nil(t, err)
	usuario, _ := usuario_entity.NovoUsuario("roberto", "123")
	err = repository.Criar(usuario)
	assert.Nil(t, err)

	var usuarioEncontrado usuario_entity.Usuario
	err = db.First(&usuarioEncontrado, "nome = ?", usuario.Nome).Error
	assert.Nil(t, err)
	assert.Equal(t, usuarioEncontrado.Nome, usuario.Nome)
	assert.Equal(t, usuarioEncontrado.Ativo, true)
	assert.Equal(t, usuarioEncontrado.Id, usuario.Id)
	assert.NotEmpty(t, usuarioEncontrado.Senha)
}

func TestBuscarUsuarioPorId(t *testing.T) {
	_, repository, err := Repository()
	assert.Nil(t, err)
	usuario, _ := usuario_entity.NovoUsuario("roberto", "123")
	_ = repository.Criar(usuario)
	usuarioEncontrado, err := repository.BuscarPorID(usuario.Id.String())
	assert.Nil(t, err)
	assert.NotNil(t, usuarioEncontrado)
	assert.Equal(t, usuarioEncontrado.Nome, usuario.Nome)
	assert.Equal(t, usuarioEncontrado.Id, usuario.Id)
}

func TestBuscarUsuario_QueNaoExiste(t *testing.T) {
	_, repository, err := Repository()
	assert.Nil(t, err)
	usuarioEncontrado, err := repository.BuscarPorID(entity.NewID().String())
	assert.Nil(t, usuarioEncontrado)
	assert.NotNil(t, err)
}

func TestBuscarTodos(t *testing.T) {
	_, repository, err := Repository()
	assert.Nil(t, err)
	usuario1, _ := usuario_entity.NovoUsuario("elton", "123")
	usuario2, _ := usuario_entity.NovoUsuario("roberto", "123")
	usuario3, _ := usuario_entity.NovoUsuario("daniel", "123")

	_ = repository.Criar(usuario1)
	_ = repository.Criar(usuario2)
	_ = repository.Criar(usuario3)

	usuarios, err := repository.BuscarTodos()
	assert.Nil(t, err)
	assert.NotNil(t, usuarios)
	assert.Equal(t, len(usuarios), 3)
}

func TestBuscarTodos_NaoEncotrado(t *testing.T) {
	_, repository, err := Repository()
	assert.Nil(t, err)
	usuarios, err := repository.BuscarTodos()
	assert.Nil(t, err)
	assert.NotNil(t, usuarios)
	assert.Equal(t, len(usuarios), 0)
}

func TestAtualizarUsuario(t *testing.T) {
	_, repository, err := Repository()
	assert.Nil(t, err)

	usuario, _ := usuario_entity.NovoUsuario("Elton", "123")
	_ = repository.Criar(usuario)

	uEncontrado, _ := repository.BuscarPorID(usuario.Id.String())
	uEncontrado.Nome = "Casacio"

	repository.Atualizar(uEncontrado)
	usuarioAtualizado, err := repository.BuscarPorID(usuario.Id.String())
	assert.Nil(t, err)
	assert.NotNil(t, usuarioAtualizado)
	assert.Equal(t, usuarioAtualizado.Nome, "Casacio")
}

func TestAtualizarUsuario_UsuarioInvalido(t *testing.T) {
	_, repository, err := Repository()
	assert.Nil(t, err)
	err = repository.Atualizar(&usuario_entity.Usuario{})
	assert.NotNil(t, err)
}

func TestExcluirUsuario(t *testing.T) {
	_, repository, err := Repository()
	assert.Nil(t, err)

	usuario, _ := usuario_entity.NovoUsuario("Elton", "123")
	_ = repository.Criar(usuario)

	repository.Excluir(usuario.Id.String())
	assert.Nil(t, err)

	uEncontrado, err := repository.BuscarPorID(usuario.Id.String())
	assert.Nil(t, err)
	assert.NotNil(t, uEncontrado)
	assert.Equal(t, uEncontrado.Nome, usuario.Nome)
	assert.Equal(t, uEncontrado.Ativo, false)
}

func TestExcluirUsuario_ID_Invalido(t *testing.T) {
	_, repository, err := Repository()
	assert.Nil(t, err)
	err = repository.Excluir(entity.NewID().String())
	assert.NotNil(t, err)
}
