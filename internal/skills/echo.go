package skills

import (
	"context"
	"fmt"
)

// EchoSkill returns the input untouched.
type EchoSkill struct{}

func (EchoSkill) Name() string {
	return "echo"
}

func (EchoSkill) Description() string {
	return "Echoes the input back to the caller"
}

func (EchoSkill) Run(_ context.Context, input string) (string, error) {
	return fmt.Sprintf("echo: %s", input), nil
}
