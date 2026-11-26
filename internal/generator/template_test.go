package generator

import (
	"strings"
	"testing"

	"github.com/hrnnsx/goutline/internal/config"
)

func TestRenderTemplate(t *testing.T) {
	cfg := config.Config{
		Project: config.ProjectConfig{
			Name:   "my-api",
			Module: "github.com/example/my-api",
		},
	}

	t.Run("successful rendering", func(t *testing.T) {
		result, err := RenderTemplate(
			"templates/base/main.go.tmpl",
			cfg,
		)

		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		output := string(result)

		expected := []string{
			"my-api",
			"github.com/example/my-api",
		}

		for _, value := range expected {
			if !strings.Contains(output, value) {
				t.Errorf("expected output to contain %q, got:\n%s", value, output)
			}
		}

		if strings.Contains(output, "{{.Project.Name}}") ||
			strings.Contains(output, "{{.Project.Module}}") {
			t.Error("template placeholders were not replaced")
		}
	})

	t.Run("template not found", func(t *testing.T) {
		_, err := RenderTemplate(
			"templates/base/not-found.go.tmpl",
			cfg,
		)

		if err == nil {
			t.Fatal("expected an error for missing template")
		}
	})

	t.Run("invalid template syntax", func(t *testing.T) {
		_, err := RenderTemplate(
			"templates/base/invalid.go.tmpl",
			cfg,
		)

		if err == nil {
			t.Fatal("expected an error for invalid template syntax")
		}
	})
}
