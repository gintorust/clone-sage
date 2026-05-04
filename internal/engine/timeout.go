package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/gintorust/clone-sage/internal/model"
)

func RunWithTimeout(check model.Check, timeoutMs int) model.Result{
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutMs))
	defer cancel()

	resultChan := make(chan model.Result, 1)

	go func(){
		resultChan <- check.Diagnose(ctx)
	}()

	select{
	case res:= <-resultChan:
		return res
	case <-ctx.Done():
		return model.Result{
			Name:     check.Name(),
			Status:   "failed",
			Severity: "blocker",
			Evidence: fmt.Sprintf("Check timed out after %dms", timeoutMs),
			Why:      "The underlying command or network request hung.",
		}
	}

}