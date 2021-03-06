package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"

	"github.com/omdbRestAPI/app/core"
	"github.com/omdbRestAPI/app/helper"
	"github.com/omdbRestAPI/app/transport/rest/v1"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		helper.Log(helper.LogTypeFatal, "Error loading .env file")
	}

	core.InitDB()
	defer core.GetDB().Close()

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		r := fmt.Sprintf("%s%s Not Found", c.Request.Host, c.Request.URL)
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": r, "success": false})
	})

	v1.Route(router)

	port := os.Getenv("PORT")
	router.Run(port)

}
