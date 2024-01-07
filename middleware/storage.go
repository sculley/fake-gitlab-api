package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/internal/storage"
)

func Storage(s *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("storage", s)
		c.Next()
	}
}
