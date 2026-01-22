package agent

import "fmt"

// ErrSkillNotFound is returned when a skill referenced in a plan is missing.
type ErrSkillNotFound struct {
	SkillName string
}

func (err ErrSkillNotFound) Error() string {
	return fmt.Sprintf("skill not found: %s", err.SkillName)
}
