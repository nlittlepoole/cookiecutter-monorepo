package controller

import (
	"github.com/labstack/echo"
	"github.com/nlittlepoole/cookiecutter-monorepo/guessing/game/constants"
	"github.com/stretchr/testify/assert"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPrompt(t *testing.T) {
	// Setup
	e := echo.New()
	e.Renderer = &Renderer{
		Templates: template.Must(template.ParseGlob("../" + constants.VIEWS_TEMPLATES)),
	}

	req := httptest.NewRequest(echo.GET, "/prompt/?key=0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	api := Controller{NewFixtureBackend()}
	api.Prepopulate()

	// Assertions
	if assert.NoError(t, api.GetPrompt(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Is it an animal?")
	}
}
