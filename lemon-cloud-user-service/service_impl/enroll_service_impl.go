package service_impl

import (
	"context"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common/dto"
)

type EnrollServiceImpl struct{}

func (EnrollServiceImpl) SignInByNumber(context.Context, *dto.SignInByNumberReq) (*dto.SignInRsp, error) {
	panic("implement me")
}
