package service

import "fops/pkg/exception"

type SystemService struct {
	Preload []func()
}

var service *SystemService

func InitialSystemService() error {
	service = &SystemService{}
	return nil
}

func GetSystemService() *SystemService {
	return service
}

func (s *SystemService) AddPreload(load func()) error {
	s.Preload = append(s.Preload, load)
	return nil
}

func (s *SystemService) Shutdown() {
	defer exception.PanicHandle()
	for _, load := range s.Preload {
		load()
	}
}
