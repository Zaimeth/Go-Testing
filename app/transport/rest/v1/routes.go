package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/omdbRestAPI/app/transport/rest/v1/delivery"
)

func Route(route *gin.Engine) *gin.RouterGroup {
	api := route.Group("/v1")
	{
		api.GET("/movies", delivery.GetMovies)
		api.GET("/movies/:id", delivery.GetDetailMovies)
	}

	return api
}
