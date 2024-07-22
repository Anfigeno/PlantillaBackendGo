package controladores

import (
	"fmt"
	"net/http"
)

var Inicio = Controlador{
	Get: inicioGet,
}

func inicioGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hola")
}
