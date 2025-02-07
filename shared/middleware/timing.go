package middleware

import (
	"context"
	"fmt"
	"shared/core"
	"time"
)

func Timing[R any, S any](actionHandler core.ActionHandler[R, S], label string) core.ActionHandler[R, S] {
	return func(ctx context.Context, request R) (*S, error) {
		start := time.Now()

		response, err := actionHandler(ctx, request)

		duration := time.Since(start)
		fmt.Printf("Request %s took %v\n", label, duration)

		return response, err
	}
}
