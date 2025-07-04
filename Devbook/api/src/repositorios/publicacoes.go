package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Publicacoes repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicacao no banco de dados
func (repositorios Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorios.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarPorID traz uma única publicação do banco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id where p.id = ?`, publicacaoID)
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

// Buscar traz as publicações dos usuários seguidos e também do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT DISTINCT p.*, u.nick 
FROM publicacoes p
JOIN usuarios u ON u.id = p.autor_id
LEFT JOIN seguidores s ON s.usuario_id = p.autor_id
WHERE p.autor_id = ? OR s.seguidor_id = ?
ORDER BY p.id DESC`,
		usuarioID, usuarioID)
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

// AtualizarPublicacao altera os dados de uma publicação no banco de dados
func (repositorio Publicacoes) AtualizarPublicacao(publicacaoId uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacoes where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}
	fmt.Println("Deletou")

	return nil
}

// BuscarPorUsuario traz todas as publicações de um usuário específico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioId uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select p.*, u.nick from publicacoes p
	inner join usuarios u on p.autor_id = u.id
	where p.autor_id = ?
	order by 1 desc`, usuarioId)
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
	statement, erro := repositorio.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Descurtir subtrai uma curtida na publicação
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
	update publicacoes set curtidas = 
	CASE WHEN curtidas > 0 THEN curtidas - 1
	ELSE 0 END where id = ?`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}
