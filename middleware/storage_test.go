package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/internal/storage"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/middleware"
)

func TestStorageMiddlware(t *testing.T) {
	s := storage.New()

	router := gin.Default()
	router.Use(middleware.Storage(s))

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authenticated"})
	})

	t.Run("Storage", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/test", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		require.Equal(t, http.StatusOK, recorder.Code)
	})
}
