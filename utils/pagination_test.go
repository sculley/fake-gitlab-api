package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/utils"
)

func TestPagination(t *testing.T) {
	router := gin.New()

	// Set up a sample route that uses the Pagination function
	router.GET("/test", func(c *gin.Context) {
		utils.Pagination(c, 20, 5)
	})

	// Create a mock HTTP request
	request := httptest.NewRequest("GET", "/test", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	require.Equal(t, "1", recorder.Header().Get("x-page"))
}
