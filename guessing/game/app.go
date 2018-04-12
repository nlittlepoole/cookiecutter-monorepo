package main

import (
	"flag"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/nlittlepoole/cookiecutter-monorepo/guessing/game/constants"
	"github.com/nlittlepoole/cookiecutter-monorepo/guessing/game/controller"
	"github.com/sandalwing/echo-logrusmiddleware"
	"html/template"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	port := flag.Int("port", 80, "The Port to serve the game from")
	flag.Parse()

	e := echo.New()
	e.Logger = logrusmiddleware.Logger{logrus.StandardLogger()}
	e.Use(logrusmiddleware.Hook())

	e.Renderer = &controller.Renderer{
		Templates: template.Must(template.ParseGlob(constants.VIEWS_TEMPLATES)),
	}
	e.Static("/static", "static/")

	api := controller.Controller{controller.NewFixtureBackend()}
	api.Prepopulate()

	e.GET("/", api.Home)
	e.GET("/prompt", api.GetPrompt)
	e.GET("/success", api.Success)
	e.GET("/add", api.AddForm)
	e.GET("/inject", api.InjectPrompt)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
