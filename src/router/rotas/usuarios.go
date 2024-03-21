package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

// Essas rotas vão lidar com o carregamento de páginas
var rotasUsuarios = []Rota{
	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
}
