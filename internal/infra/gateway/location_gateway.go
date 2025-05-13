package gateway

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/israelalvesmelo/desafio-temperature/cmd/api/config"
	"github.com/israelalvesmelo/desafio-temperature/internal/domain/dto"
	"github.com/israelalvesmelo/desafio-temperature/internal/domain/entity"
	gatewaydomain "github.com/israelalvesmelo/desafio-temperature/internal/domain/gateway"
)

type LocationGatewayImpl struct {
	config *config.CEP
}

func NewLocationGateway(config *config.CEP) gatewaydomain.LocationGateway {
	return LocationGatewayImpl{
		config: config,
	}
}

var createCepEndpoint = func(baseUrl, cep string) string {
	return strings.Join([]string{baseUrl, "ws", cep, "json"}, "/")
}

func (g LocationGatewayImpl) GetLocation(ctx context.Context, cep string) (*entity.Location, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req, reqErr := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		createCepEndpoint(g.config.URL, cep),
		nil,
	)
	if reqErr != nil {
		return nil, reqErr
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, doErr := client.Do(req)
	if doErr != nil {
		return nil, doErr
	}
	defer resp.Body.Close()

	bodyBytes, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	var location dto.LocationOutput
	if unmErr := json.Unmarshal(bodyBytes, &location); unmErr != nil {
		return nil, unmErr
	}

	if location.HasError() {
		return nil, entity.ErrZipcodeNotFound
	}

	return &entity.Location{
		Localidade: location.Localidade,
	}, nil
}
