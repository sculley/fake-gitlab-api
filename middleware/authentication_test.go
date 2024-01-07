package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/middleware"
)

func TestAuthenticationMiddleware(t *testing.T) {
	router := gin.Default()
	router.Use(middleware.Authentication())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authenticated"})
	})

	t.Run("Authenticated", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/test", nil)
		request.Header.Set("PRIVATE-TOKEN", "valid_token")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		require.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("Not Authenticated", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/test", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		require.Equal(t, http.StatusUnauthorized, recorder.Code)
	})
}
