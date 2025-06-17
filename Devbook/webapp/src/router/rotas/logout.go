package rotas

import (
	"webapp/src/controllers"
)

var rotaLogout = Rota{
	URI:                "/logout",
	Metodo:             "GET",
	Funcao:             controllers.FazerLogout,
	RequerAutenticacao: true,
}
