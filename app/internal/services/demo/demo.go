package demo

import (
	"github.com/prf16/go-zero-box-rpc/app/internal/model/usermodel"
)

type Service struct {
	UserModel usermodel.UserModel
}

func NewService(userModel usermodel.UserModel) *Service {
	return &Service{UserModel: userModel}
}
