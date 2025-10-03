package gomockbug

import (
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

type MyType struct{}

func (MyType) MyMethod(param1 *time.Time, param2 string) {}

type MyInterface interface {
	MyMethod(param1 *time.Time, param2 string)
}

//go:generate mockgen -typed -destination=mock.go -package=gomockbug . MyInterface

func TestFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	mock := NewMockMyInterface(ctrl)
	mock.EXPECT().MyMethod(&now, "right")

	mock.MyMethod(&now, "wrong")
}

func TestPanics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockMyInterface(ctrl)
	mock.EXPECT().MyMethod(nil, "right")

	mock.MyMethod(nil, "wrong")
}
