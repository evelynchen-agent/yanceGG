package agent

import (
	"context"
	"errors"
)

// Skill represents a capability the agent can execute.
type Skill interface {
	Name() string
	Description() string
	Run(ctx context.Context, input string) (string, error)
}

// SkillSet stores the registered skills by name.
type SkillSet struct {
	registry map[string]Skill
}

// NewSkillSet initializes an empty skill set.
func NewSkillSet() SkillSet {
	return SkillSet{registry: make(map[string]Skill)}
}

// Register adds a skill to the set.
func (s SkillSet) Register(skill Skill) error {
	if skill == nil {
		return errors.New("skill is nil")
	}
	if s.registry == nil {
		s.registry = make(map[string]Skill)
	}
	s.registry[skill.Name()] = skill
	return nil
}

// Get returns a skill by name.
func (s SkillSet) Get(name string) (Skill, bool) {
	skill, ok := s.registry[name]
	return skill, ok
}

// All returns a copy of the skills.
func (s SkillSet) All() []Skill {
	skills := make([]Skill, 0, len(s.registry))
	for _, skill := range s.registry {
		skills = append(skills, skill)
	}
	return skills
}
