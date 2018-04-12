package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddYesPrompt(t *testing.T) {
	p := Prompt{Key: "0", Value: "Are you organic?"}
	p.NewYesKey()
	assert.Equal(t, "01", p.YesKey, "YesKey should be Key + 1")
}

func TestAddNoPrompt(t *testing.T) {
	p := Prompt{Key: "0", Value: "Are you organic?"}
	p.NewNoKey()
	assert.Equal(t, "00", p.NoKey, "NoKey should be Key + 0")
}
