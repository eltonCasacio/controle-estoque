package mysql_repository

import (
	"database/sql"

	"github.com/eltonCasacio/controle-estoque/internal/domain/usuario/entity"
	usuario "github.com/eltonCasacio/controle-estoque/internal/infrastructure/usuario/model"
)

type UsuarioRepository struct {
	DB *sql.DB
}

func NovoUsuarioRpository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

func (u *UsuarioRepository) Criar(usuario *entity.Usuario) error {
	stmt, err := u.DB.Prepare("insert into usuarios(id, nome, senha, ativo) values(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		usuario.GetID(),
		usuario.GetNome(),
		usuario.GetSenha(),
		usuario.IsAtivo(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *UsuarioRepository) BuscarPorID(id string) (*entity.Usuario, error) {
	stmt, err := u.DB.Prepare("select * from usuarios where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var usuarioModel usuario.UsuarioModel
	err = stmt.QueryRow(id).Scan(&usuarioModel.Id, &usuarioModel.Nome, &usuarioModel.Senha, &usuarioModel.Ativo)
	if err != nil {
		return nil, err
	}

	var usuario entity.Usuario
	usuario.ChangeID(usuarioModel.Id)
	usuario.ChangeNome(usuarioModel.Nome)
	usuario.ChangeSenha(usuarioModel.Senha)
	usuario.Desativar()
	if usuarioModel.Ativo {
		usuario.Ativar()
	}
	return &usuario, nil
}

func (u *UsuarioRepository) BuscarTodos() ([]entity.Usuario, error) {
	rows, err := u.DB.Query("select * from usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []entity.Usuario

	for rows.Next() {
		var usuario usuario.UsuarioModel
		err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Senha, &usuario.Ativo)
		if err != nil {
			return nil, err
		}
		var u entity.Usuario
		u.ChangeID(usuario.Id)
		u.ChangeNome(usuario.Nome)
		u.ChangeSenha(usuario.Senha)
		u.Desativar()
		if usuario.Ativo {
			u.Ativar()
		}

		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}

func (u *UsuarioRepository) Atualizar(usuario *entity.Usuario) error {
	stmt, err := u.DB.Prepare("update usuarios set id = ?, nome = ?, senha = ?, ativo = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(usuario.GetID(), usuario.GetNome(), usuario.GetSenha(), usuario.IsAtivo())
	if err != nil {
		return err
	}
	return nil
}

func (u *UsuarioRepository) Excluir(id string) error {
	stmt, err := u.DB.Prepare("delete from usuarios where id = ?")
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

func (u *UsuarioRepository) BuscarPorNome(nome string) (*entity.Usuario, error) {
	stmt, err := u.DB.Prepare("select * from usuarios where nome = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var usuarioModel usuario.UsuarioModel
	err = stmt.QueryRow(nome).Scan(&usuarioModel.Id, &usuarioModel.Nome, &usuarioModel.Senha, &usuarioModel.Ativo)
	if err != nil {
		return nil, err
	}

	var usuario entity.Usuario
	usuario.ChangeID(usuarioModel.Id)
	usuario.ChangeNome(usuarioModel.Nome)
	usuario.ChangeSenha(usuarioModel.Senha)
	usuario.Desativar()
	if usuarioModel.Ativo {
		usuario.Ativar()
	}
	return &usuario, nil
}

func (u *UsuarioRepository) BuscarPaginado(page, limit string, sort string) ([]entity.Usuario, error) {
	rows, err := u.DB.Query("SELECT * from usuarios LIMIT ?, ?;", page, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []entity.Usuario

	for rows.Next() {
		var usuario usuario.UsuarioModel
		err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Senha, &usuario.Ativo)
		if err != nil {
			return nil, err
		}
		var u entity.Usuario
		u.ChangeID(usuario.Id)
		u.ChangeNome(usuario.Nome)
		u.ChangeSenha(usuario.Senha)
		u.Desativar()
		if usuario.Ativo {
			u.Ativar()
		}

		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}
