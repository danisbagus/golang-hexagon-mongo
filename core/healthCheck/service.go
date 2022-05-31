package healthcheck

import (
	port "github.com/danisbagus/golang-hexagon-mongo/core/port/healthCheck"
)

type Service struct {
	repo port.Repository
}

func New(repo port.Repository) port.Service {
	return &Service{repo: repo}
}

func (s *Service) Ping() map[string]string {
	if err := s.repo.Ping(); err != nil {
		return map[string]string{"status": "NOT OK"}
	}
	return map[string]string{"status": "OK"}
}
