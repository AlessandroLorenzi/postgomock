package mypackage

import "github.com/AlessandroLorenzi/postgomock/xyz"

type Service struct {
	xyzSvc *xyz.XYZ
}

func New(xyzSvc *xyz.XYZ) *Service {
	return &Service{xyzSvc: xyzSvc}
}

func (s *Service) GetInfoFromXyz() (string, error) {
	data, err := s.xyzSvc.GetInfo(42)
	if err != nil {
		return "", err
	}
	return data.InfoINeed, nil
}
