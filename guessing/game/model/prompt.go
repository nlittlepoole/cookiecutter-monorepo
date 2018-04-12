/*
Package model implements data structures for guessing game
*/
package model

// Prompt is an object representing a Question or Answer in the Game
type Prompt struct {
	Key    string `json:"key" form:"key" query:"key"`
	Value  string `json:"value" form:"value" query:"value"`
	YesKey string
	NoKey  string
}

// NewYesKey generates the corresponding key for the pointer to the
// Prompt corresponding to a "yes" answer
func (p *Prompt) NewYesKey() {
	p.YesKey = p.Key + "1"
}

// NewNoKey generates the corresponding key for the pointer to the
// Prompt corresponding to a "no" answer
func (p *Prompt) NewNoKey() {
	p.NoKey = p.Key + "0"
}
