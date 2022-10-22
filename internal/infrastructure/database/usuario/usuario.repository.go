package database

import (
	e "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func NovoUsuarioRpository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

func (u *UsuarioRepository) Criar(usuario *e.Usuario) error {
	return u.DB.Create(usuario).Error
}

func (u *UsuarioRepository) BuscarPorID(id entity.ID) (*e.Usuario, error) {
	var usuario e.Usuario
	if err := u.DB.Where("id = ?", id).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (u *UsuarioRepository) BuscarTodos() ([]e.Usuario, error) {
	var usuarios []e.Usuario
	if err := u.DB.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (u *UsuarioRepository) Atualizar(usuario *e.Usuario) error {
	_, err := u.BuscarPorID(usuario.Id)
	if err != nil {
		return err
	}
	return u.DB.Save(usuario).Error
}

func (u *UsuarioRepository) Excluir(id entity.ID) error {
	usuario, err := u.BuscarPorID(id)
	if err != nil {
		return err
	}
	usuario.Ativo = false
	return u.DB.Save(usuario).Error
}
