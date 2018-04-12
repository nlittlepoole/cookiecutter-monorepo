package controller

import(
	"io"
	"html/template"
	"net/http"
	"github.com/labstack/echo"
	"github.com/nlittlepoole/cookiecutter-monorepo/guessing/game/model"
)

type Backend interface {
	GetPrompt(p *model.Prompt) (*model.Prompt, error)
	SetPrompt(p *model.Prompt) (Backend, error)
	InjectPrompt(p *model.Prompt, answer string) (Backend, error)
	Values() []model.Prompt
}

type Controller struct {
	Backend Backend
}

func (c *Controller) Prepopulate() error {
	root := &model.Prompt{
		Key: "0",
		Value: "Is it an animal?",
		YesKey: "",
		NoKey: "",
	}
	root.NewYesKey()
	root.NewNoKey()

	left := &model.Prompt{
		Key: root.NoKey,
		Value: "You are thinking of a computer!",
	}

	right := &model.Prompt{
		Key: root.YesKey,
		Value: "You are thinking of a dog!",
	}

	c.Backend, _ = c.Backend.SetPrompt(root)
	c.Backend, _ = c.Backend.SetPrompt(left)
	c.Backend, _ = c.Backend.SetPrompt(right)
	return nil
}

func (c *Controller) GetPrompt(context echo.Context) (err error) {
	p := new(model.Prompt)
	if err = context.Bind(p); err != nil {
	    return err
	}
	p, err = c.Backend.GetPrompt(p)
	if err != nil {
		return err
	}
    return context.Render(http.StatusOK, "prompt", p)
}

func (c *Controller) InjectPrompt(context echo.Context) (err error) {
	p := &model.Prompt{}
	answer := context.QueryParam("answer")
	if err = context.Bind(p); err != nil {
	    return err
	}
	c.Backend, err = c.Backend.InjectPrompt(p, answer)
	if err != nil {
		return err
	}
    return c.Home(context)
}

func (c *Controller) AddForm(context echo.Context) (err error) {
	p := new(model.Prompt)
	if err = context.Bind(p); err != nil {
	    return err
	}
    return context.Render(http.StatusOK, "add", p)
}

func (c *Controller) Home (context echo.Context) (err error) {
	return context.Render(http.StatusOK, "home", "")
}

func (c *Controller) Success (context echo.Context) (err error) {
	return context.Render(http.StatusOK, "success", "")
}

type Renderer struct {
    templates *template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return r.templates.ExecuteTemplate(w, name, data)
}