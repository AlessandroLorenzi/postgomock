package anotherpackage

import (
	"fmt"

	"github.com/AlessandroLorenzi/postgomock/xyz"
)

type Service struct {
	xyzSvc *xyz.XYZ
}

func New(xyzSvc *xyz.XYZ) *Service {
	return &Service{xyzSvc: xyzSvc}
}

func (s *Service) DoStuff(info string) {
	fmt.Println(info)
}
