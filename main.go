package main

import (
	"os"

	"github.com/AlessandroLorenzi/postgomock/anotherpackage"
	"github.com/AlessandroLorenzi/postgomock/mypackage"
	"github.com/AlessandroLorenzi/postgomock/xyz"
)

func main() {
	xyzSvc := xyz.New(
		os.Getenv("XYZ_KEY_ID"),
		os.Getenv("XYZ_SECRET_KEY"),
	)

	myPackageSvc := mypackage.New(xyzSvc)
	anotherSvc := anotherpackage.New(xyzSvc)

	theValue, _ := myPackageSvc.GetInfoFromXyz()

	anotherSvc.DoStuff(theValue)
}
