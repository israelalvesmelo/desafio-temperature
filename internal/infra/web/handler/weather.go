package handler

import (
	"net/http"

	"github.com/israelalvesmelo/desafio-temperature/internal/domain/usecase"
)

type WeatherHandler struct {
	useCase *usecase.GetWeatherUseCase
}

func NewWeatherHandler(useCase *usecase.GetWeatherUseCase) *WeatherHandler {
	return &WeatherHandler{
		useCase: useCase,
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {

}
