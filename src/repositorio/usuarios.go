package repositorio

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositoriosDeUsuarios cria uma estancia nova de usuarios
func NovoRepositoriosDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("Insert Into usuarios(nome,nick,email,senha) values (?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	ultimoID, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoID), nil
}

// Buscar traz todos os usuarios baseados em nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	linhas, erro := repositorio.db.Query("Select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u where nome like ? or nick like ?", nomeOuNick, nomeOuNick)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarPorId traz um registro baseado em ID
func (repositorio Usuarios) BuscarPorID(usuarioID uint64) (modelos.Usuario, error) {

	linhas, erro := repositorio.db.Query("Select id,nome, nick, email, criadoEm from usuarios where id = ? ", usuarioID)

	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(usuarioID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome =?, nick=?, email = ? where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuarioID); erro != nil {
		return erro
	}
	return nil
}

// Deletar remove um usuario do banco de dados
func (repositorio Usuarios) Deletar(usuarioId uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id=?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioId); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna o seu id e senha com hash
func (repositorio Usuarios) BuscaPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("Select id,senha from usuarios where email = ?", email)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Seguir permite um usuário seguir outro
func (repositorio Usuarios) Seguir(usuarioId, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("Insert ignore into seguidores(usuario_id, seguidor_id) values(?,?)")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, seguidorID); erro != nil {
		return erro
	}
	return nil
}

// PararDeSeguir permite um usuário seguir outro
func (repositorio Usuarios) PararDeSeguir(usuarioId, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("Delete from seguidores where usuario_id = ? and seguidor_id=?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// BuscarSeguidores traz lista de seguidores
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		Select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u
		inner join seguidores s on u.id = s.seguidor_id 
		where s.usuario_id = ?
	`, usuarioID)

	if erro != nil {
		return []modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarSeguindo traz lista de seguidores
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		Select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u
		inner join seguidores s on u.id = s.usuario_id 
		where s.seguidor_id = ?
	`, usuarioID)

	if erro != nil {
		return []modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
