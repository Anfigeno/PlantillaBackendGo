package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func FormatoDeRegistros() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func CargarVariablesDeEntorno() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error al cargar las variables de entorno")
	}
}
