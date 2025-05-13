package main

import (
	"fmt"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/israelalvesmelo/desafio-temperature/cmd/api/config"
	"github.com/israelalvesmelo/desafio-temperature/internal/domain/usecase"
	"github.com/israelalvesmelo/desafio-temperature/internal/infra/gateway"
	"github.com/israelalvesmelo/desafio-temperature/internal/infra/web/handler"
	"github.com/israelalvesmelo/desafio-temperature/internal/infra/web/webserver"
)

func main() {
	// Load config
	var cfg config.Config
	viperCfg := config.NewViper("../../env")
	viperCfg.ReadViper(&cfg)

	// Create gateway
	locationGateway := gateway.NewLocationGateway(&cfg.CEP)
	temperatureGateway := gateway.NewTemperatureGateway(&cfg.Temperature)

	// Create use case
	useCase := usecase.NewGetTemperatureUseCase(locationGateway, temperatureGateway)

	// Create handler
	handler := handler.NewTemperatureHandler(useCase)

	// Create server
	server := webserver.NewWebServer(fmt.Sprintf(":%s", cfg.App.Port))
	server.AddMiddleware(chimiddleware.Logger)

	server.AddHandler("/temperature", handler.GetWeather)

	server.Start()

}
