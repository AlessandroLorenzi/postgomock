package mypackage

import (
	"testing"

	"github.com/AlessandroLorenzi/postgomock/xyz"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GetInfoFromXyz(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedXyx := NewMockXyz(mockCtrl)
	mockedXyx.EXPECT().GetInfo(42).Return(&xyz.GetInfoOutput{
		InfoINeed: "something",
	}, nil)

	s := New(mockedXyx)

	out, err := s.GetInfoFromXyz()

	assert.Nil(t, err)
	assert.Equal(t, "something", out)

}
