package service

import "fops/expection"

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

func (s *SystemService) Shutdown () {
	defer expection.PanicHandle()
	for _, load := range s.Preload {
		load()
	}
}
