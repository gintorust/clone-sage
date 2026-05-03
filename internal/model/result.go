package model

import "time"

// Result represents the outcome of a single diagnostic check.
type Result struct {
	Name     string        `json:"name"`
	Type     string        `json:"type"`
	Status   string        `json:"status"`   // e.g., "passed", "failed"
	Severity string        `json:"severity"` // e.g., "info", "warning", "blocker"
	Title    string        `json:"title,omitempty"`
	Evidence string        `json:"evidence"`
	Why      string        `json:"why,omitempty"`
	Fix      string        `json:"fix,omitempty"`
	Duration time.Duration `json:"duration,omitempty"` // Track slow checks
}