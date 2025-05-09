package gateway

import (
	"context"

	"github.com/israelalvesmelo/desafio-temperature/internal/domain/entity"
)

type LocationGateway interface {
	GetLocation(ctx context.Context, cep string) (*entity.Location, error)
}

type TemperatureGateway interface {
	GetTempCelsius(ctx context.Context, location string) (*float64, error)
}
