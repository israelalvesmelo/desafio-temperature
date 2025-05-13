package usecase

import (
	"context"
	"fmt"

	"github.com/israelalvesmelo/desafio-temperature/internal/domain/dto"
	"github.com/israelalvesmelo/desafio-temperature/internal/domain/entity"
	"github.com/israelalvesmelo/desafio-temperature/internal/domain/gateway"
)

type GetTemperatureUseCase struct {
	cepGateway         gateway.LocationGateway
	temperatureGateway gateway.TemperatureGateway
}

func NewGetTemperatureUseCase(
	cepGateway gateway.LocationGateway,
	temperatureGateway gateway.TemperatureGateway,
) *GetTemperatureUseCase {
	return &GetTemperatureUseCase{
		cepGateway:         cepGateway,
		temperatureGateway: temperatureGateway,
	}
}

func (uc *GetTemperatureUseCase) Execute(ctx context.Context, cep string) (*dto.TemperatureOutput, error) {
	location, err := uc.cepGateway.GetLocation(ctx, cep)
	if err != nil {
		return nil, err
	}

	fmt.Println("location: ", *location)

	temperatureCelsius, err := uc.temperatureGateway.GetTempCelsius(ctx, location.Localidade)
	if err != nil {
		return nil, err
	}
	fmt.Println("temperature celsius: ", *temperatureCelsius)
	temperature := entity.NewTemperature(*temperatureCelsius)

	return &dto.TemperatureOutput{
		TempC: temperature.Celsius(),
		TempF: temperature.Fahrenheit(),
		TempK: temperature.Kelvin(),
	}, nil

}
