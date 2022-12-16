package rotas

import (
	"api/src/controller"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controller.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controller.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controller.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controller.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controller.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/seguindo",
		Metodo:             http.MethodGet,
		Funcao:             controller.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controller.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
