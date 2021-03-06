package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omdbRestAPI/app/helper"
	"github.com/omdbRestAPI/app/transport/rest/v1/models"
	"github.com/omdbRestAPI/app/transport/rest/v1/usecase"
)

func GetMovies(c *gin.Context) {
	var (
		omdbData   models.OMDBSerializerResponses
		queryParam models.QueryParam
		err        error
	)

	if err = c.Bind(&queryParam); err != nil {
		errResp := &helper.ErrorResp{Error: err, Code: http.StatusNotFound, Message: "Failed to parsing data"}
		errResp.HandleError(c)
		return
	}

	if omdbData, err = usecase.GetMovies(queryParam); err != nil {
		errResp := &helper.ErrorResp{Error: err, Code: http.StatusNotFound, Message: "Failed get movies"}
		errResp.HandleError(c)
		return
	}

	totalPage, _ := strconv.Atoi(omdbData.TotalResults)

	c.JSON(http.StatusOK, gin.H{
		"data":    omdbData.OmdbResponses,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"meta": models.Meta{
			TotalPage:   int32(totalPage / 10),
			TotalRecord: int32(totalPage),
			Limit:       10,
			PageNumber:  queryParam.Page,
		},
	})
}

func GetDetailMovies(c *gin.Context) {
	var (
		omdbData   models.OMDBSerializerResponse
		queryParam models.QueryParam
		err        error
	)

	queryParam.ID = c.Params.ByName("id")

	if omdbData, err = usecase.GetDetailMovies(queryParam); err != nil {
		errResp := &helper.ErrorResp{Error: err, Code: http.StatusNotFound, Message: "Failed get detail movies"}
		errResp.HandleError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    omdbData,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}
