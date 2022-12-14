package rotas

import (
	"api/src/controller"
	"net/http"
)

var rotaLogin = Rota{
	Uri:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controller.Login,
	RequerAutenticacao: false,
}
