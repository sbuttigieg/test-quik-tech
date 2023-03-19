package api

import (
	"time"

	"github.com/google/uuid"
	service "github.com/sbuttigieg/test-quik-tech/internal/services/api"
)

func NewService(uuidFunc func() uuid.UUID, timeFunc func() time.Time) (service.Service, error) {
	service := service.New(uuidFunc, timeFunc)

	return service, nil
}
