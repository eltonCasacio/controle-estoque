package usuario_repository

import (
	usuario_entity "github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	"github.com/eltonCasacio/controle-estoque/pkg/entity"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	DB *gorm.DB
}

func NovoUsuarioRpository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

func (u *UsuarioRepository) Create(usuario *usuario_entity.Usuario) error {
	return u.DB.Create(usuario).Error
}

func (u *UsuarioRepository) Find(id entity.ID) (*usuario_entity.Usuario, error) {
	var usuario usuario_entity.Usuario
	if err := u.DB.Where("id = ?", id).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (u *UsuarioRepository) FindAll() (*[]usuario_entity.Usuario, error) {
	var usuarios []usuario_entity.Usuario
	if err := u.DB.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return &usuarios, nil
}

func (u *UsuarioRepository) Update(usuario *usuario_entity.Usuario) error {
	if err := u.DB.Updates(&usuario).Error; err != nil {
		return err
	}
	return nil
}

func (u *UsuarioRepository) Delete(id entity.ID) error {
	if err := u.DB.Where("id = ?", id).UpdateColumn("ativo", id).Error; err != nil {
		return err
	}
	return nil
}
