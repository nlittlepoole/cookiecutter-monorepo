package main

import (
	"html/template"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/sandalwing/echo-logrusmiddleware"
	"github.com/nlittlepoole/cookiecutter-monorepo/guessing/game/controller"
	"github.com/nlittlepoole/cookiecutter-monorepo/guessing/game/constants"
)



func init(){
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()
	e.Logger = logrusmiddleware.Logger{logrus.StandardLogger()}
	e.Use(logrusmiddleware.Hook())

	e.Renderer =  &controller.Renderer{
	    templates: template.Must(template.ParseGlob(constants.VIEWS_TEMPLATES)),
	}
	e.Static("/static", "static/")

	api := controller.Controller{controller.NewFixtureBackend()}
	api.Prepopulate()

	e.GET("/", api.Home)
	e.GET("/prompt", api.GetPrompt)
	e.GET("/success", api.Success)
	e.GET("/add", api.AddForm)
	e.GET("/inject", api.InjectPrompt)
	e.Logger.Fatal(e.Start(":80"))
}