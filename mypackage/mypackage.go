package mypackage

import (
	"github.com/AlessandroLorenzi/postgomock/xyz"
	_ "github.com/golang/mock/mockgen/model"
)

type Service struct {
	xyzSvc Xyz
}

//go:generate mockgen -destination=./mock_xyz.go -package=mypackage . Xyz
type Xyz interface {
	GetInfo(int) (*xyz.GetInfoOutput, error)
}

func New(xyzSvc Xyz) *Service {
	return &Service{xyzSvc: xyzSvc}
}

func (s *Service) GetInfoFromXyz() (string, error) {
	data, err := s.xyzSvc.GetInfo(42)
	if err != nil {
		return "", err
	}
	return data.InfoINeed, nil
}
