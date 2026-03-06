package model

import (
	"github.com/google/wire"
	"github.com/prf16/go-zero-box-rpc/app/internal/model/messagemodel"
	"github.com/prf16/go-zero-box-rpc/app/internal/model/usermodel"
)

var Provider = wire.NewSet(
	NewModel,
	usermodel.NewUserModel,
	messagemodel.NewMessageModel,
)

type Model struct {
	UserModel    usermodel.UserModel
	MessageModel messagemodel.MessageModel
}

func NewModel(userModel usermodel.UserModel, messageModel messagemodel.MessageModel) *Model {
	return &Model{UserModel: userModel, MessageModel: messageModel}
}
