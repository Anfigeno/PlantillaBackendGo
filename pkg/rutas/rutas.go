package rutas

import (
	"net/http"

	"github.com/Anfigeno/PlantillaBackendGo/pkg/controladores"
)

func RegistrarRutas(r *http.ServeMux) {
	r.HandleFunc("GET /", controladores.Inicio.Get)
}
