package model

import (
	"context"
)

// Check is the strict contract that all diagnostic implementations must follow.
type Check interface {
	// Name returns the unique name of the check (e.g., "redis-reachable")
	Name() string
	
	// Type returns the rule type (e.g., "tcp_reachable")
	Type() string
	
	// Diagnose executes the actual logic. The context is passed so the engine 
	// can forcibly time it out if it hangs.
	Diagnose(ctx context.Context) Result
	
}