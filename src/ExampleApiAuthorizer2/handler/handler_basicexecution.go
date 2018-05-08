package handler

import (
	"context"
	"fmt"
)

func BasicExecutionEventHandler(ctx context.Context, event interface{}) (string, error) {
	return fmt.Sprintf("Hello, world."), nil
}
