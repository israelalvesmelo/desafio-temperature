package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/israelalvesmelo/desafio-temperature/internal/domain/usecase"
)

type TemperatureHandler struct {
	useCase *usecase.GetTemperatureUseCase
}

func NewTemperatureHandler(useCase *usecase.GetTemperatureUseCase) *TemperatureHandler {
	return &TemperatureHandler{
		useCase: useCase,
	}
}

func (h *TemperatureHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		fmt.Println("params CEP is required")
		Error(w, "params CEP is required", http.StatusBadRequest)
		return
	}

	weather, err := h.useCase.Execute(r.Context(), cep)
	if err != nil {
		fmt.Println("Error:", err)
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(weather); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
