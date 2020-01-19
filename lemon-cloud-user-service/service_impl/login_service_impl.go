package service_impl

import (
	"context"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common/dto"
)

type LoginServiceImpl struct{}

func (LoginServiceImpl) LoginByNumber(context.Context, *dto.LoginByNumberReq) (*dto.LoginRsp, error) {
	return &dto.LoginRsp{
		ServiceInfo: nil,
		Token:       "123456",
		ExpiredAt:   654321,
	}, nil
}
