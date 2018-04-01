package controller

import(
	"sync"
	"errors"
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
	fmt.Println("%s", c.Backend.Values())
	return nil
}

func (c *Controller) GetPrompt(c echo.Context) (err error) {
	p := new(model.Prompt)
	if err = c.Bind(p); err != nil {
	    return err
	}
	fmt.Println(*p)
	p, err = c.Backend.GetPrompt(p)
	if err != nil {
		return err
	}
    return c.Render(http.StatusOK, "prompt", p)
}

func (c *Controller) InjectPrompt(c echo.Context) (err error) {
	p := &model.Prompt{}
	answer := c.QueryParam("answer")
	if err = c.Bind(p); err != nil {
	    return err
	}
	c.Backend, err = c.Backend.InjectPrompt(p, answer)
	if err != nil {
		return err
	}
    return c.Home(c)
}

func (c *Controller) AddForm(c echo.Context) (err error) {
	p := new(model.Prompt)
	if err = c.Bind(p); err != nil {
	    return err
	}
    return c.Render(http.StatusOK, "add", p)
}

func (c *Controller) Home (c echo.Context) (err error) {
	return c.Render(http.StatusOK, "home", "")
}

func (c *Controller) Success (c echo.Context) (err error) {
	return c.Render(http.StatusOK, "success", "")
}


type FixtureBackend struct {
	values map[string]model.Prompt
	mutex *sync.Mutex
}

func (f FixtureBackend) GetPrompt(p *model.Prompt) (*model.Prompt, error) {
	if val, ok := f.values[p.Key]; ok {
	    return &val, nil
	} else {
		return nil, errors.New("Prompt with id " + p.Key + " not found")
	}
}

func (f FixtureBackend) SetPrompt(p *model.Prompt) (Backend, error) {
	f.mutex.Lock()
	f.values[p.Key] = *p
	f.mutex.Unlock()
	return f, nil
}

func (f FixtureBackend) SetFixturePrompt(p *model.Prompt) (FixtureBackend, error) {
	f.mutex.Lock()
	f.values[p.Key] = *p
	f.mutex.Unlock()
	return f, nil
}

func (f FixtureBackend) InjectPrompt(p *model.Prompt, answer string) (Backend, error) {
	p.NewYesKey()
	p.NewNoKey()

	left, _ := f.GetPrompt(p)
	left.Key = p.NoKey

	right := &model.Prompt{
		Key: p.YesKey,
		Value: answer,
	}

	f, _ = f.SetFixturePrompt(p)
	f, _ = f.SetFixturePrompt(left)
	f, _ = f.SetFixturePrompt(right)
	return f, nil
}

func (f FixtureBackend) Values() []model.Prompt {
	values := make([]model.Prompt, 0)
	for _, v := range f.values {
		values = append(values, v)
	}
	return values
}

func NewFixtureBackend() FixtureBackend {
	mapping := make(map[string]model.Prompt)
	return FixtureBackend{
		values: mapping,
		mutex: &sync.Mutex{},
	}
}

func (f FixtureBackend) SetFixturePrompt(p *Prompt) (FixtureBackend, error) {
	f.mutex.Lock()
	f.values[p.Key] = *p
	f.mutex.Unlock()
	return f, nil
}

func (f FixtureBackend) InjectPrompt(p *Prompt, answer string) (Backend, error) {
	p.NewYesKey()
	p.NewNoKey()

	left, _ := f.GetPrompt(p)
	left.Key = p.NoKey

	right := &Prompt{
		Key: p.YesKey,
		Value: answer,
	}

	f, _ = f.SetFixturePrompt(p)
	f, _ = f.SetFixturePrompt(left)
	f, _ = f.SetFixturePrompt(right)
	return f, nil
}

func (f FixtureBackend) Values() []Prompt {
	values := make([]Prompt, 0)
	for _, v := range f.values {
		values = append(values, v)
	}
	return values
}

func NewFixtureBackend() FixtureBackend {
	mapping := make(map[string]Prompt)
	return FixtureBackend{
		values: mapping,
		mutex: &sync.Mutex{},
	}
}