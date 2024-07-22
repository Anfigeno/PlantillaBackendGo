package servidor

import (
	"net/http"
	"os"

	"github.com/Anfigeno/PlantillaBackendGo/pkg/rutas"
	"github.com/rs/zerolog/log"
)

func Iniciar() {
	r := http.NewServeMux()

	rutas.RegistrarRutas(r)

	go iniciarServidorHttp(r)
}

func iniciarServidorHttp(r *http.ServeMux) {
	puerto := os.Getenv("PUERTO")

	err := http.ListenAndServe(":"+puerto, r)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error al iniciar el servidor http")
	}
}
