package api

import (
	"context"
)

func (s *service) Credit(ctx context.Context, param1, param2 string) ([]string, error) {

	return []string{param1, param2}, nil
}
