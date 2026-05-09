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
Evaluate and interpret considering general tennis statistics - if a value is low (e.g. 1 unforced error), do NOT describe it as a problem.
Prefer neutral or positive phrasing when metrics are within good ranges.

%s
%s
%s
%s
%s
%s
%s
%s
%s

Input:
First serve in (percent): %d
Second serve in (percent): %d
Unforced errors (in total): %d
`,
		addMeaningfulIssueRequirement(),
		addValidJSONRequirement(),
		addSeverityRequirement(),
		addStatsInterpretationRequirement(),
		addAntiHallucinationRequirement(),
		addFocusAreaEnumerationRequirement(),
		addRecommendationRequirement(),
		addLimitRequirement(),
		addExampleOutputRequirement(),
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

%s
%s
%s
%s
%s
%s
%s
%s

Input: %s
`,
		addValidJSONRequirement(),
		addMeaningfulIssueRequirement(),
		addSeverityRequirement(),
		addAntiHallucinationRequirement(),
		addFocusAreaEnumerationRequirement(),
		addRecommendationRequirement(),
		addLimitRequirement(),
		addExampleOutputRequirement(),
		i.Text,
	)
}

func addValidJSONRequirement() string {
	return `
Output must be valid JSON only.

Return ONLY valid JSON:

{
  "issues": [
    {
      "text": "string",
      "severity": "low"
    }
  ],
  "recommendations": [
    "string"
  ],
  "focus_area": "string"
}

Do not wrap JSON in markdown.
Do not use markdown fence.
No explanations.
No extra text.

Each issue must include:
text (string)
severity (one of: low, medium, high)
`
}

func addMeaningfulIssueRequirement() string {
	return `
It is acceptable to return empty issue list:
"issues": []

If statistics are within normal ranges,
do not create artificial weaknesses.

Do not criticize isolated low values
unless they clearly indicate a performance problem.

Issues must describe concrete weaknesses, not just symptoms.
`
}

func addSeverityRequirement() string {
	return `
Severity rules:
high: directly impacts performance or consistency
medium: noticeable but not critical
low: minor inefficiencies or observations
	`
}

func addStatsInterpretationRequirement() string {
	return `
General interpretation guidelines:

First serve in percentage:
- below 50 = poor consistency
- 50-65 = average
- above 65 = strong consistency

Second serve in percentage:
- below 75 = unreliable
- 75-85 = acceptable
- above 85 = solid and reliable

Unforced errors:
- high only if clearly excessive
- isolated or low counts should not be criticized

Winners vs unforced errors:
- significantly more unforced errors than winners may indicate overaggression or inconsistency

Return performance:
- low return percentages may indicate positioning or timing issues

Double faults:
- repeated double faults may indicate second serve instability
`
}

func addAntiHallucinationRequirement() string {
	return `
Do not infer:
- emotions
- confidence
- mentality
- fatigue
- focus
- motivation
- psychology

unless explicitly supported by the input.

Issues must reference observable match data only.
`
}

func addFocusAreaEnumerationRequirement() string {
	return `
focus_area must be exactly one of:

serve
return
consistency
movement
aggression
net_play
mental
fitness
`
}

func addRecommendationRequirement() string {
	return `
Recommendations must be:
- specific
- actionable
- tennis-related

Avoid generic advice.
Recommend specific exercises or drills if possible.
Recommendations must directly relate
to detected issues or provided match data.

Bad recommendations:
- "Practice more"
- "Improve consistency"
- "Focus harder"

Good recommendations:
- "Practice second serve placement under pressure"
- "Use cross-court rally drills to reduce unforced forehand errors"
- "Train return timing against wide serves"
`
}

func addLimitRequirement() string {
	return `
Return:
- maximum 3 issues
- maximum 5 recommendations	
`
}

func addExampleOutputRequirement() string {
	return `
Example valid response:

{
  "issues": [
    {
      "text": "Second serve consistency drops under pressure",
      "severity": "medium"
    }
  ],
  "recommendations": [
    "Practice second serve placement under pressure situations",
    "Use basket drills focused on kick serve consistency"
  ],
  "focus_area": "serve"
}
`
}
