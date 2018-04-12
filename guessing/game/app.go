// Binary libgame accepts port flags, builds server and runs game serving static assets
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

// init instantiates the logrus logformatter in JSON format
func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

// main parses command line arguments, instantiates the server, and sets routes
func main() {
	port := flag.Int("port", 80, "The Port to serve the game from")
	flag.Parse()

	e := echo.New()
	middleware := logrusmiddleware.Logger{}
	middleware.Logger = logrus.StandardLogger()
	e.Logger = middleware
	e.Use(logrusmiddleware.Hook())

	e.Renderer = &controller.Renderer{
		Templates: template.Must(template.ParseGlob(constants.ViewsTemplates)),
	}
	e.Static("/static", "static/")

	api := controller.Controller{Backend: controller.NewFixtureBackend()}
	api.Prepopulate()

	e.GET("/", api.Home)
	e.GET("/prompt", api.GetPrompt)
	e.GET("/success", api.Success)
	e.GET("/add", api.AddForm)
	e.GET("/inject", api.InjectPrompt)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
