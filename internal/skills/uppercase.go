package skills

import (
	"context"
	"strings"
)

// UppercaseSkill capitalizes the input.
type UppercaseSkill struct{}

func (UppercaseSkill) Name() string {
	return "uppercase"
}

func (UppercaseSkill) Description() string {
	return "Transforms the input into upper-case"
}

func (UppercaseSkill) Run(_ context.Context, input string) (string, error) {
	return strings.ToUpper(input), nil
}
