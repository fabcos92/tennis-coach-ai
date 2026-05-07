package llm

import (
	"fmt"
	"tennis-coach-ai/internal/application/ports"
	"tennis-coach-ai/internal/domain/input"
)

type DefaultPromptBuilder struct{}

func NewDefaultPromptBuilder() ports.PromptBuilder {
	return &DefaultPromptBuilder{}
}

func (b *DefaultPromptBuilder) BuildStats(i *input.Stats) string {
	return fmt.Sprintf(`You are an expert tennis coach.

You MUST base all conclusions strictly on provided numerical data.

Rules:
Do NOT exaggerate or generalize.
Only describe something as "high", "low", or "poor" if the data clearly supports it.
Evaulate and interpret considering general tennis statistics - if a value is low (e.g. 1 unforced error), do NOT describe it as a problem.
Prefer neutral or positive phrasing when metrics are within good ranges.
Do NOT invent issues.
Recommend specific exercises or drills if possible.

Output must be valid JSON only.

Return issues as structured objects.

Each issue must include:
text (string)
severity (one of: low, medium, high)

Severity rules:
high: directly impacts performance or consistency
medium: noticeable but not critical
low: minor inefficiencies or observations

Return ONLY JSON:
{
  "issues": [{"",""}, {"",""}],
  "recommendations": ["","",...],
  "focus_area": ""
}

No markdown.
No explanations.
No extra text.
Issues must describe concrete weaknesses, not just symptoms.
Each field must be filled.
Input stats:
First serve in (percent): %d
Second serve in (percent): %d
Unforced errors (in total): %d
`,
		i.FirstServe.In,
		i.SecondServe.In,
		i.UnforcedErrors,
	)
}

func (b *DefaultPromptBuilder) BuildText(i *input.Text) string {
	return fmt.Sprintf(`You are an expert tennis coach.

You analyze a textual description of a match.

Rules:
Base conclusions only on information present in the text.
Do NOT assume statistics that are not mentioned.
Do NOT exaggerate problems.
Highlight key patterns, strengths, and weaknesses.
If no clear issues are present, return an empty "issues" array.
Recommend specific exercises or drills if possible.

Output must be valid JSON only.
Return issues as structured objects.

Each issue must include:
text (string)
severity (one of: low, medium, high)

Severity rules:
high: directly impacts performance or consistency
medium: noticeable but not critical
low: minor inefficiencies or observations

Return ONLY JSON:
{
  "issues": [{"",""}, {"",""}],
  "recommendations": ["","",...],
  "focus_area": ""
}

No markdown.
No explanations.
No extra text.
Issues must describe concrete weaknesses, not just symptoms.
Each field must be filled.
Input: %s
`,
		i.Text,
	)
}
