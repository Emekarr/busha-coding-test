package movie

import (
	"encoding/json"
	"errors"

	"github.com/emekarr/coding-test-busha/logger"
	"github.com/emekarr/coding-test-busha/network"
	"go.uber.org/zap"
)

var MovieService = movieService{network: network.NetworkController{BaseUrl: "https://swapi.dev/api/films"}}

type movieService struct {
	network network.NetworkController
}

func (ms *movieService) FetchMovies() (*[]Movie, error) {
	response := ms.network.Get("/", nil, nil)
	if response.Error != nil {
		logger.Error(errors.New("error fetching movies from movie service"), zap.Error(response.Error))
		return nil, response.Error
	}
	var parsedResponse MovieResponse
	json.Unmarshal(*response.Body, &parsedResponse)
	return &parsedResponse.Results, nil
}