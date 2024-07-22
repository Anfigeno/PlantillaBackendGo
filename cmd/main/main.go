package main

import (
	"github.com/Anfigeno/PlantillaBackendGo/pkg/config"
	"github.com/Anfigeno/PlantillaBackendGo/pkg/servidor"
	"github.com/Anfigeno/PlantillaBackendGo/pkg/util"
	"github.com/rs/zerolog/log"
)

func init() {
	config.FormatoDeRegistros()
	config.CargarVariablesDeEntorno()
}
func main() {
	log.Info().Msg("Iniciando servidor!")
	servidor.Iniciar()
	util.CongelarHastaRecibirSeñal()
}
