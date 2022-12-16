package repositorio

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

// NovoRepositoriosDePublicacoes cria uma estancia nova de repositorio para publicacoes
func NovoRepositoriosDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into publicacao (titulo,conteudo, autor_id) values(?,?,?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)

	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// BuscarPorID busca uma unica publicacao do banco
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
			select p.*, u.nick from publicacao p
			inner join usuarios u on u.id = p.autor_id
			where p.id = ?
		`, publicacaoID)

	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil

}

// Buscar retorna as publicações do usuario logado e dos que ele segue
func (repositorio Publicacoes) Buscar(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
					select distinct p.*, u.nick from publicacao p
					inner join usuarios u on u.id = p.autor_id
					inner join seguidores s on p.autor_id = s.usuario_id 
					where u.id = ? or s.seguidor_id = ?
					order by 1 desc 
					`, usuarioId, usuarioId)

	if erro != nil {
		return []modelos.Publicacao{}, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

// Atualizar retorna as publicações do usuario logado e dos que ele segue
func (repositorio Publicacoes) Atualizar(publicacaoId uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacao set titulo =?, conteudo =? where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação no banco
func (repositorio Publicacoes) Deletar(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacao where id=?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}
	return nil
}

// BuscarPorUsuario traz uma publicação por usuario
func (repositorio Publicacoes) BuscarPorUsuario(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick from publicacao p 
		join usuarios u on u.id = p.autor_id
		where p.autor_id = ?
	`, usuarioId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()
	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

// Curtir adiciona uma curtida na publicação
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("update publicacao set curtidas = curtidas+1 where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Descurtir remove uma curtida no banco
func (repositorio Publicacoes) DesCurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
						update publicacao set curtidas = 
						Case when curtidas > 0 then curtidas -1 
						Else curtidas end
						where id = ?`)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}
