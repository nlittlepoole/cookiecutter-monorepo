package main

import (
	"io"
	"html/template"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/sandalwing/echo-logrusmiddleware"
	"github.com/nlittlepoole/cookiecutter-monorepo/guessing/game/controller"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func init(){
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()
	e.Logger = logrusmiddleware.Logger{logrus.StandardLogger()}
	e.Use(logrusmiddleware.Hook())

	t := &Template{
	    templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t
	e.Static("/static", "static/")

	api := controller.Api{model.NewFixtureBackend()}
	api.Prepopulate()

	e.GET("/", api.Home)
	e.GET("/prompt", api.GetPrompt)
	e.GET("/success", api.Success)
	e.GET("/add", api.AddForm)
	e.GET("/inject", api.InjectPrompt)
	e.Logger.Fatal(e.Start(":80"))
}