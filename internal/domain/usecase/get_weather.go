package usecase

import (
	"context"

	"github.com/israelalvesmelo/desafio-temperature/internal/domain/dto"
	"github.com/israelalvesmelo/desafio-temperature/internal/domain/entity"
	"github.com/israelalvesmelo/desafio-temperature/internal/domain/gateway"
)

type GetWeatherUseCase struct {
	cepGateway         gateway.LocationGateway
	temperatureGateway gateway.TemperatureGateway
}

func NewGetWeatherUseCase(
	cepGateway gateway.LocationGateway,
	temperatureGateway gateway.TemperatureGateway,
) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		cepGateway:         cepGateway,
		temperatureGateway: temperatureGateway,
	}
}

func (uc *GetWeatherUseCase) Execute(ctx context.Context, cep string) (*dto.TemperatureOutput, error) {
	location, err := uc.cepGateway.GetLocation(ctx, cep)
	if err != nil {
		return nil, err
	}

	temperatureCelsius, err := uc.temperatureGateway.GetTempCelsius(ctx, location.Localidade)
	if err != nil {
		return nil, err
	}
	temperature := entity.NewTemperature(*temperatureCelsius)

	return &dto.TemperatureOutput{
		Location: location.Localidade,
		TempC:    temperature.Celsius(),
		TempF:    temperature.Fahrenheit(),
		TempK:    temperature.Kelvin(),
	}, nil

}
