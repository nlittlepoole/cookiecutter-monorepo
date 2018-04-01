package main 

import(
	"github.com/nlittlepoole/cookiecutter-monorepo/guessing/game/model"
)

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