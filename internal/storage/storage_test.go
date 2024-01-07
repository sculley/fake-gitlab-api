package storage_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/internal/storage"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/models"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TestSetAndGet(t *testing.T) {
	templateFile := "/Users/sculley/Go/src/gitlab.com/someadmin/go/components/fake-gitlab-api/data/json/get_projects_response.json.tmpl"

	name := utils.GenerateRandomWord()

	response, err := utils.RenderTemplate(templateFile, map[string]interface{}{
		"id":           utils.GenerateRandomID(),
		"name":         cases.Title(language.English).String(name),
		"path":         name,
		"description":  "Example Project",
		"namespace_id": utils.GenerateRandomID(),
		"namespace":    utils.GenerateRandomWord(),
		"url":          "http://localhost:8080",
		"timestamp":    utils.GenerateTimestamp(),
	})
	if err != nil {
		t.Error(err)
		return
	}

	s := storage.New()

	s.AppendProject(response)

	projects := s.GetProjects()

	var jsonResponse []models.Project
	err = json.Unmarshal(projects[0], &jsonResponse)
	if err != nil {
		t.Error(err)
		return
	}

	require.NotEmpty(t, projects)
	require.NoError(t, err)
}
