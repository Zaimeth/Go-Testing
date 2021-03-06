package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/omdbRestAPI/app/transport/rest/v1/delivery"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("v1/movies", delivery.GetMovies)

	err := godotenv.Load("../../.env")

	if err != nil {
		t.Fatal("Failed to load env ", err)
	}

	req, err := http.NewRequest(http.MethodGet, os.Getenv("HOST")+os.Getenv("PORT")+"/v1/movies?search=batman&page=7", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestGetDetailMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()

	r.GET("v1/movies", delivery.GetDetailMovies)

	err := godotenv.Load("../../.env")

	if err != nil {
		t.Fatal("Failed to load env ", err)
	}

	req, err := http.NewRequest(http.MethodGet, os.Getenv("HOST")+os.Getenv("PORT")+"/v1/movies/tt12794046", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
