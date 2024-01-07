package utils_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/someadmin/go/components/fake-gitlab-api/utils"
)

func TestCommonUtils(t *testing.T) {
	t.Run("RenderTempate", func(t *testing.T) {
		templateFile, err := os.CreateTemp("", "")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(templateFile.Name())

		_, err = templateFile.WriteString("Hello {{ .name }}")
		if err != nil {
			t.Fatal(err)
		}

		response, err := utils.RenderTemplate(templateFile.Name(), map[string]interface{}{
			"name": "World",
		})
		if err != nil {
			t.Fatal(err)
		}

		require.Equal(t, "Hello World", string(response))
		require.NoError(t, err)
	})

	t.Run("RenderTemplateOpenFileError", func(t *testing.T) {
		_, err := utils.RenderTemplate("non-existing-file", nil)
		require.Error(t, err)
	})

	t.Run("RenderTemplateParseError", func(t *testing.T) {
		templateFile, err := os.CreateTemp("", "")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(templateFile.Name())

		_, err = templateFile.WriteString("Hello {{ .name ")
		if err != nil {
			t.Fatal(err)
		}

		_, err = utils.RenderTemplate(templateFile.Name(), nil)
		require.Error(t, err)
	})

	t.Run("GenerateRandomID", func(t *testing.T) {
		id := utils.GenerateRandomID()
		require.NotEmpty(t, id)
	})

	t.Run("GenerateTimestamp", func(t *testing.T) {
		ts := utils.GenerateTimestamp()
		require.NotEmpty(t, ts)
	})

	t.Run("GenerateRandomWord", func(t *testing.T) {
		word := utils.GenerateRandomWord()
		require.NotEmpty(t, word)
	})
}
