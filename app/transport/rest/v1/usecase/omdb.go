package usecase

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/omdbRestAPI/app/transport/rest/v1/models"
	"github.com/omdbRestAPI/app/transport/rest/v1/services"
	"os"
)

func GetMovies(queryParam models.QueryParam) (models.OMDBSerializerResponses, error) {
	var (
		responseOmdb models.OMDBSerializerResponses
		responseMap  map[string]interface{}
		err          error
	)

	path := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", os.Getenv("OMDB_API_URL"), os.Getenv("OMDB_API_KEY"), queryParam.Search, queryParam.Page)

	if responseMap, err = services.GetMoviesByParam(path); err != nil {
		return responseOmdb, err
	}

	if err := mapstructure.Decode(responseMap, &responseOmdb); err != nil {
		return responseOmdb, err
	}

	return responseOmdb, nil
}

func GetDetailMovies(queryParam models.QueryParam) (models.OMDBSerializerResponse, error) {
	var (
		responseOmdb models.OMDBSerializerResponse
		responseMap  map[string]interface{}
		err          error
	)

	path := fmt.Sprintf("%s?apikey=%s&i=%s", os.Getenv("OMDB_API_URL"), os.Getenv("OMDB_API_KEY"), queryParam.ID)

	if responseMap, err = services.GetMoviesByParam(path); err != nil {
		return responseOmdb, err
	}

	if err := mapstructure.Decode(responseMap, &responseOmdb); err != nil {
		return responseOmdb, err
	}

	return responseOmdb, nil
}
