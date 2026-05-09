package analysis

import "errors"

var (
	ErrInvalidAnalysis             = errors.New("invalid analysis")
	ErrInvalidFocusArea            = errors.New("invalid focus_area")
	ErrInvalidSeverity             = errors.New("invalid severity")
	ErrEmptyIssueText              = errors.New("empty issue text")
	ErrIssuesSizeExceeded          = errors.New("issues size exceeded")
	ErrRecommendationsSizeExceeded = errors.New("recommendations size exceeded")
)
