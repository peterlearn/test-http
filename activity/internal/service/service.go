package service

import (
	"github.com/google/wire"
	activity "x-proto/activity"
)

var ProviderSet = wire.NewSet(New)

// Service service.
type Service struct {
	activity.UnimplementedActivityServer
}

// New new a service and return.
func New() (s *Service, cf func(), err error) {
	s = &Service{}
	cf = func() {}

	return
}
