package model

type Prompt struct {
	Key    string `json:"key" form:"key" query:"key"`
	Value  string `json:"value" form:"value" query:"value"`
	YesKey string
	NoKey  string
}

func (p *Prompt) NewYesKey() {
	p.YesKey = p.Key + "1"
}

func (p *Prompt) NewNoKey() {
	p.NoKey = p.Key + "0"
}
