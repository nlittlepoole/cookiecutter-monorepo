/*
Package controller contains all logic connecting models to views
*/
package controller

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/nlittlepoole/cookiecutter-monorepo/guessing/game/model"
	"html/template"
	"io"
	"net/http"
	"sync"
)

// Backend interface represents a storage backend for the Question Tree
type Backend interface {
	GetPrompt(p *model.Prompt) (*model.Prompt, error)
	SetPrompt(p *model.Prompt) (Backend, error)
	InjectPrompt(p *model.Prompt, answer string) (Backend, error)
	Values() []model.Prompt
}

// Controller is an struct class with methods that use the Backend
// to connect models to views
type Controller struct {
	Backend Backend
}

// Prepopulate seeds the backend with a starting question and two answers
func (c *Controller) Prepopulate() error {
	root := &model.Prompt{
		Key:    "0",
		Value:  "Is it an animal?",
		YesKey: "",
		NoKey:  "",
	}
	root.NewYesKey()
	root.NewNoKey()

	left := &model.Prompt{
		Key:   root.NoKey,
		Value: "You are thinking of a computer!",
	}

	right := &model.Prompt{
		Key:   root.YesKey,
		Value: "You are thinking of a dog!",
	}

	c.Backend, _ = c.Backend.SetPrompt(root)
	c.Backend, _ = c.Backend.SetPrompt(left)
	c.Backend, _ = c.Backend.SetPrompt(right)
	return nil
}

// GetPrompt loads the rest of the struct data from the backend, based on the key
// provided to the object
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

// InjectPrompt utilizes the Backend's InjectPrompt logic based on request parameters
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

// AddForm Renders the "add question" form, based on the provided Prompt
func (c *Controller) AddForm(context echo.Context) (err error) {
	p := new(model.Prompt)
	if err = context.Bind(p); err != nil {
		return err
	}
	return context.Render(http.StatusOK, "add", p)
}

// Home renders the home page of the game
func (c *Controller) Home(context echo.Context) (err error) {
	return context.Render(http.StatusOK, "home", "")
}

// Success renders a page signaling the user got the correct answer
func (c *Controller) Success(context echo.Context) (err error) {
	return context.Render(http.StatusOK, "success", "")
}

// Renderer is an object implementing logic to render View Templates
type Renderer struct {
	Templates *template.Template
}

// Render is a method required by Echo in order to convert interfaces to HTML
func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.Templates.ExecuteTemplate(w, name, data)
}

// FixtureBackend is an in memory, map powered backend to test with
type FixtureBackend struct {
	values map[string]model.Prompt
	mutex  *sync.Mutex
}

// GetPrompt returns a pointer to the prompt in the map with the same key as provided
func (f FixtureBackend) GetPrompt(p *model.Prompt) (*model.Prompt, error) {
	if val, ok := f.values[p.Key]; ok {
		return &val, nil
	}
	return nil, errors.New("Prompt with id " + p.Key + " not found")
}

// SetPrompt locks the map, adds a mapping to the pointer of the given object
// based on key, and then returns FixtureBackend itself in Interface form
func (f FixtureBackend) SetPrompt(p *model.Prompt) (Backend, error) {
	f.mutex.Lock()
	f.values[p.Key] = *p
	f.mutex.Unlock()
	return f, nil
}

// setFixturePrompt  does the same as SetPrompt but returns the FixtureBackend in its default type
func (f FixtureBackend) setFixturePrompt(p *model.Prompt) (FixtureBackend, error) {
	f.mutex.Lock()
	f.values[p.Key] = *p
	f.mutex.Unlock()
	return f, nil
}

// InjectPrompt creates a new prompt with the given Key
// it moves the old Prompt with that key to the No Child slot
// Finally it creates a 3rd prompt, with answer as a value, that occupies the Yes slot
// Intiutively this is the same as injecting a new question into the tree
func (f FixtureBackend) InjectPrompt(p *model.Prompt, answer string) (Backend, error) {
	p.NewYesKey()
	p.NewNoKey()

	left, _ := f.GetPrompt(p)
	left.Key = p.NoKey

	right := &model.Prompt{
		Key:   p.YesKey,
		Value: answer,
	}

	f, _ = f.setFixturePrompt(p)
	f, _ = f.setFixturePrompt(left)
	f, _ = f.setFixturePrompt(right)
	return f, nil
}

// Values returns a slice of all the Prompts in the tree
func (f FixtureBackend) Values() []model.Prompt {
	values := make([]model.Prompt, 0)
	for _, v := range f.values {
		values = append(values, v)
	}
	return values
}

// NewFixtureBackend allocates a map and mutex for FixtureBackend and
// initiates the struct
func NewFixtureBackend() FixtureBackend {
	mapping := make(map[string]model.Prompt)
	return FixtureBackend{
		values: mapping,
		mutex:  &sync.Mutex{},
	}
}
