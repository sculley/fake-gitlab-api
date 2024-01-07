package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/controllers"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/internal/storage"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/middleware"
)

func main() {
	router := gin.Default()

	s := storage.New()

	router.Use(middleware.Authentication())
	router.Use(middleware.Storage(s))

	router.GET("/api/v4/projects", controllers.ProjectsControllerCreate)
	router.POST("/api/v4/projects/", controllers.ProjectsControllerCreate)

	router.Run(":8080")
}
