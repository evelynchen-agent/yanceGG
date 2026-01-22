package agent

import "context"

// Agent ties together a planner role and a set of skills.
type Agent struct {
	Planner PlanRole
	Skills  SkillSet
}

// BuildPlan creates a plan for a goal using the planner role.
func (a Agent) BuildPlan(ctx context.Context, goal string) (Plan, error) {
	skills := a.Skills.All()
	return a.Planner.Plan(ctx, goal, skills)
}

// ExecutePlan runs the skills specified by a plan in order.
func (a Agent) ExecutePlan(ctx context.Context, plan Plan) ([]Result, error) {
	results := make([]Result, 0, len(plan.Steps))
	for _, step := range plan.Steps {
		skill, ok := a.Skills.Get(step.SkillName)
		if !ok {
			return results, ErrSkillNotFound{SkillName: step.SkillName}
		}

		result, err := skill.Run(ctx, step.Input)
		if err != nil {
			return results, err
		}

		results = append(results, Result{
			Step:   step,
			Output: result,
		})
	}

	return results, nil
}
