package services

import (
	"fmt"
	"tennis-coach-ai/internal/models"
)

func buildPrompt(req models.AnalyzeRequest) string {
	base := `You are an expert tennis coach.

Return ONLY JSON:
{
  "issues": [],
  "recommendations": [],
  "focus_area": ""
}

You must return ONLY valid JSON.
No markdown.
No explanations.
No extra text.
Issues must describe concrete weaknesses, not just symptoms.
Each field must be filled.
`

	if req.Type == "match_stats" {
		return base + fmt.Sprintf(`
Input stats:
First serve: %d
Second serve won: %d
Unforced errors: %d
`,
			req.Stats.FirstServePct,
			req.Stats.SecondServeWonPct,
			req.Stats.UnforcedErrors,
		)
	}

	return base + "\nInput:\n" + req.Text
}
