package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/internal/storage"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/models"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/utils"
)

func ProjectsControllerCreate(c *gin.Context) {
	s := c.MustGet("storage").(*storage.Storage)

	if c.Request.Method == "GET" {
		response := s.GetProjects()

		var jsonResponse []models.Project
		for _, project := range response {
			var projectResponse models.Project
			err := json.Unmarshal(project, &projectResponse)
			if err != nil {
				c.String(http.StatusInternalServerError, "Internal Server Error")
				return
			}

			jsonResponse = append(jsonResponse, projectResponse)
		}

		postsPerPage := 20

		if c.Query("per_page") != "" {
			postsPerPage, _ = strconv.Atoi(c.Query("per_page"))
		}

		// set the pagination headers
		utils.Pagination(c,
			len(jsonResponse),
			postsPerPage,
		)

		// If the number of posts is less than or equal to the number of posts per page
		if len(jsonResponse) <= postsPerPage {
			c.JSON(200, jsonResponse)
			return
		} else {
			c.JSON(200, jsonResponse[:postsPerPage])
			return
		}
	}

	if c.Request.Method == "POST" {
		requestData := make(map[string]interface{})

		// Attempt to bind the request body to a map and return a 400 if it fails
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.String(http.StatusBadRequest, "Bad Request")
			return
		}

		// Check if name and path are missing and return a 400 if they are
		if requestData["name"] == "" && requestData["path"] == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "name, path are missing, at least one parameter must be provided",
			})
			return
		}

		templateFile := filepath.Join(".", "data", "json", "project_response.json.tmpl")

		_, err := os.Stat(templateFile)
		if os.IsNotExist(err) {
			log.Panicf("Template file %s does not exist", templateFile)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		// Parse the template and return the response
		response, err := utils.RenderTemplate(templateFile, map[string]interface{}{
			"id":           utils.GenerateRandomID(),
			"name":         requestData["name"],
			"path":         requestData["path"],
			"description":  requestData["description"],
			"namespace_id": requestData["namespace_id"],
			"namespace":    utils.GenerateRandomWord(),
			"url":          c.Request.Host,
			"timestamp":    utils.GenerateTimestamp(),
		})
		if err != nil {
			log.Printf("Error parsing template: %s", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		s.AppendProject(response)

		var jsonResponse models.Project
		err = json.Unmarshal(response, &jsonResponse)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %s", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		c.JSON(200, jsonResponse)
	}
}
