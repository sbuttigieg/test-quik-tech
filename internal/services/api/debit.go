package api

import (
	"context"
)

func (s *service) Debit(ctx context.Context, id, param1, param2 string) ([]string, error) {

	return []string{id, param1, param2}, nil
}
