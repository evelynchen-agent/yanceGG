package agent

import (
	"context"
	"fmt"
)

// PlanRole defines how to turn a goal into steps.
type PlanRole interface {
	Plan(ctx context.Context, goal string, skills []Skill) (Plan, error)
}

// BasicPlanRole is a simple planner that maps every skill to a step.
type BasicPlanRole struct{}

// Plan creates a plan that enumerates the available skills.
func (BasicPlanRole) Plan(_ context.Context, goal string, skills []Skill) (Plan, error) {
	steps := make([]Step, 0, len(skills))
	for _, skill := range skills {
		steps = append(steps, Step{
			SkillName: skill.Name(),
			Input:     fmt.Sprintf("goal: %s", goal),
		})
	}

	return Plan{
		Goal:  goal,
		Steps: steps,
	}, nil
}
