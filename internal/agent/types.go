package agent

// Plan describes the steps the agent should execute.
type Plan struct {
	Goal  string `json:"goal"`
	Steps []Step `json:"steps"`
}

// Step maps to a skill invocation.
type Step struct {
	SkillName string `json:"skill"`
	Input     string `json:"input"`
}

// Result captures the output of executing a step.
type Result struct {
	Step   Step   `json:"step"`
	Output string `json:"output"`
}
