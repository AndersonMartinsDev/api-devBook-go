package rotas

import (
	"api/src/controller"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controller.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controller.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeletarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarPublicacaoPorUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}/curtir",
		Metodo:             http.MethodPut,
		Funcao:             controller.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}/descurtir",
		Metodo:             http.MethodPut,
		Funcao:             controller.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
}
